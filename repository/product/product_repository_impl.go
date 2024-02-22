package product

import (
	"github.com/thanhvdt/vcs-week2/model"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{Db: db}
}

func (p *ProductRepositoryImpl) GetProductsWithSupplier() ([]model.Product, error) {
	var products []model.Product
	if err := p.Db.Preload("Supplier").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepositoryImpl) GetProductAboveAveragePrice() ([]model.Product, error) {
	var products []model.Product
	if err := p.Db.Where("unit_price > (?)", p.Db.Table("products").Select("AVG(unit_price)")).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
