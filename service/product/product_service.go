package product

import "github.com/thanhvdt/vcs-week2/model"

type ProductService interface {
	GetProductsWithSupplier() ([]model.Product, error)
	GetProductAboveAveragePrice() ([]model.Product, error)
	GetAllInElasticSearch() (map[string]interface{}, error)
	GetByIdInElasticSearch(docID string) (map[string]interface{}, error)
	UpdateInElasticSearch(docID string, updateFields map[string]interface{}) error
	CreateInElasticSearch(document *map[string]interface{}) (string, error)
	DeleteInElasticSearch(docID string) error
}
