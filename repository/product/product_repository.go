package product

import "github.com/thanhvdt/vcs-week2/model"

type ProductRepository interface {
	GetProductsWithSupplier() ([]model.Product, error)
	GetProductAboveAveragePrice() ([]model.Product, error)
}
