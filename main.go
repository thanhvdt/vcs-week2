package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-playground/validator/v10"
	zerolog "github.com/rs/zerolog/log"
	"github.com/thanhvdt/vcs-week2/config"
	"github.com/thanhvdt/vcs-week2/controller"
	"github.com/thanhvdt/vcs-week2/repository/category"
	"github.com/thanhvdt/vcs-week2/repository/customer"
	"github.com/thanhvdt/vcs-week2/repository/product"
	categoryservice "github.com/thanhvdt/vcs-week2/service/category"
	customerservice "github.com/thanhvdt/vcs-week2/service/customer"
	productservice "github.com/thanhvdt/vcs-week2/service/product"
	"log"
	"net/http"
	"os"
)

var es *elasticsearch.Client

// @title           Swagger Northwind APi
// @version         1.0
// @description     This is a sample server.

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	zerolog.Info().Msg("Server Started!")
	var err error

	cert, _ := os.ReadFile("/home/thanhvdt/Downloads/elasticsearch-8.12.1/config/certs/http_ca.crt")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "w*ln7FQyqlj1M*z*wpJt",
		CACert:   cert,
	}

	for {
		es, err = elasticsearch.NewClient(cfg)
		if err != nil {
			zerolog.Error().Err(err).Msg("Error creating the client")
		} else {
			break
		}
		res, _ := es.Info()

		fmt.Println(res)
	}

	// Define the query
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"public_customers_company_name": "Company",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("search-customer"),
		es.Search.WithBody(&buf),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status())
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the error message from Elasticsearch
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	// Decode the response body to a map or a specific struct
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	prettyJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %s", err)
	} else {
		fmt.Printf("Search Results: %s\n", string(prettyJSON))
	}

	db := config.ConnectDatabase()
	validate := validator.New()

	customerRepository := customer.NewCustomerRepository(db, es)
	customerService := customerservice.NewCustomerService(customerRepository, validate)
	customerController := controller.NewCustomerController(customerService)

	categoryRepository := category.NewCategoryRepository(db)
	categoryService := categoryservice.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	productRepository := product.NewProductRepository(db, es)
	productService := productservice.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)

	controllers := &Controllers{
		CustomerController: customerController,
		CategoryController: categoryController,
		ProductController:  productController,
	}
	routes := NewRouter(controllers)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err = server.ListenAndServe()
	if err != nil {
		zerolog.Fatal().Err(err).Msg("Server Stopped")
	}

}
