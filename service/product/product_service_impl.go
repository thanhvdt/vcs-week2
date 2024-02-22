package product

import (
	"github.com/go-playground/validator/v10"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/repository/product"
)

type ProductServiceImpl struct {
	productRepository product.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(productRepository product.ProductRepository, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{productRepository: productRepository, Validate: validate}
}

func (p *ProductServiceImpl) GetProductsWithSupplier() ([]model.Product, error) {
	return p.productRepository.GetProductsWithSupplier()
}

func (p *ProductServiceImpl) GetProductAboveAveragePrice() ([]model.Product, error) {
	return p.productRepository.GetProductAboveAveragePrice()
}
