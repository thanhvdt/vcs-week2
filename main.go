package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/thanhvdt/vcs-week2/config"
	"github.com/thanhvdt/vcs-week2/controller"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/repository/category"
	"github.com/thanhvdt/vcs-week2/repository/customer"
	category2 "github.com/thanhvdt/vcs-week2/service/category"
	customer2 "github.com/thanhvdt/vcs-week2/service/customer"
	"net/http"
)

func main() {
	log.Info().Msg("Server Started!")
	db := config.ConnectDatabase()
	validate := validator.New()

	customerRepository := customer.NewCustomerRepository(db)
	customerService := customer2.NewCustomerService(customerRepository, validate)
	customerController := controller.NewCustomerController(customerService)

	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category2.NewCategoryService(categoryRepository, validate)
	categoryController := controller.NewCategoryController(categoryService)

	controllers := &Controllers{
		CustomerController: customerController,
		CategoryController: categoryController,
	}
	routes := NewRouter(controllers)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Server Stopped")
	}

	var customer model.Customer
	result := db.First(&customer)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println("Connected successfully, first customer:", customer)
	}
}
