package customer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/thanhvdt/vcs-week2/model"
	"gorm.io/gorm"
	"io"
	"log"
)

type CustomerRepositoryImpl struct {
	Db *gorm.DB
	Es *elasticsearch.Client
}

func NewCustomerRepository(db *gorm.DB, es *elasticsearch.Client) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{Db: db, Es: es}
}

func (c *CustomerRepositoryImpl) Save(customer *model.Customer) (*model.Customer, error) {
	err := c.Db.Create(customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerRepositoryImpl) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := c.Db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *CustomerRepositoryImpl) FindByID(customerID string) (*model.Customer, error) {
	var customer model.Customer
	err := c.Db.Where("customer_id = ?", customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepositoryImpl) Update(customer *model.Customer) (*model.Customer, error) {
	err := c.Db.Save(customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerRepositoryImpl) Delete(customerID string) error {
	var customer model.Customer
	err := c.Db.Where("customer_id = ?", customerID).Delete(&customer).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepositoryImpl) SearchByCompany(company string) (map[string]interface{}, error) {
	// Define the query
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"public_customers_company_name": company,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	res, err := c.Es.Search(
		c.Es.Search.WithContext(context.Background()),
		c.Es.Search.WithIndex("search-customer"),
		c.Es.Search.WithBody(&buf),
		c.Es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Error closing response body: %s", err)
		}
	}(res.Body)

	fmt.Println(res.Status())
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
			return nil, err
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
		return nil, err
	}

	prettyJSON, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %s", err)
		return nil, err
	} else {
		fmt.Printf("Search Results: %s\n", string(prettyJSON))
	}

	return r, nil
}
