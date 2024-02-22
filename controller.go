package main

import "github.com/thanhvdt/vcs-week2/controller"

type Controllers struct {
	CustomerController *controller.CustomerController
	CategoryController *controller.CategoryController
	ProductController  *controller.ProductController
}
