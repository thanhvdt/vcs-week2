package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/thanhvdt/vcs-week2/config"
	"github.com/thanhvdt/vcs-week2/controller"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/repository"
	"github.com/thanhvdt/vcs-week2/service"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	log.Info().Msg("Server Started!")
	db := config.ConnectDatabase()
	validate := validator.New()

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository, validate)
	customerController := controller.NewCustomerController(customerService)
	routes := NewRouter(customerController)

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
