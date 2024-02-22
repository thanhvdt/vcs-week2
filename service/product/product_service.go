package product

import "github.com/thanhvdt/vcs-week2/model"

type ProductService interface {
	GetProductsWithSupplier() ([]model.Product, error)
	GetProductAboveAveragePrice() ([]model.Product, error)
}
