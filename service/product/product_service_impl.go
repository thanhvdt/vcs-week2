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

func (p *ProductServiceImpl) UpdateInElasticSearch(docID string, updateFields map[string]interface{}) error {
	return p.productRepository.UpdateInElasticSearch(docID, updateFields)
}

func (p *ProductServiceImpl) CreateInElasticSearch(document *map[string]interface{}) (string, error) {
	return p.productRepository.CreateInElasticSearch(document)
}

func (p *ProductServiceImpl) GetAllInElasticSearch() (map[string]interface{}, error) {
	return p.productRepository.GetAllInElasticSearch()
}

func (p *ProductServiceImpl) GetByIdInElasticSearch(docID string) (map[string]interface{}, error) {
	return p.productRepository.GetByIdInElasticSearch(docID)
}

func (p *ProductServiceImpl) DeleteInElasticSearch(docID string) error {
	return p.productRepository.DeleteInElasticSearch(docID)
}
