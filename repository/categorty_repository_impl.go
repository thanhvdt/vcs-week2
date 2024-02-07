package repository

import (
	"github.com/thanhvdt/vcs-week2/model"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{Db: db}
}

func (c *CategoryRepositoryImpl) Save(category *model.Category) (*model.Category, error) {
	err := c.Db.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) FindAll() ([]model.Category, error) {
	var categories []model.Category
	err := c.Db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryRepositoryImpl) FindByID(categoryID string) (*model.Category, error) {
	var category model.Category
	err := c.Db.Where("category_id = ?", categoryID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryRepositoryImpl) Update(category *model.Category) (*model.Category, error) {
	err := c.Db.Save(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryRepositoryImpl) Delete(categoryID string) error {
	var category model.Category
	err := c.Db.Where("category_id = ?", categoryID).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
