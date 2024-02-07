package service

import (
	"github.com/thanhvdt/vcs-week2/model"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	Db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{Db: db}
}

func (c *CategoryServiceImpl) Create(category *model.Category) (*model.Category, error) {
	err := c.Db.Create(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryServiceImpl) ReadAll() ([]model.Category, error) {
	var categories []model.Category
	err := c.Db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *CategoryServiceImpl) ReadByID(categoryID string) (*model.Category, error) {
	var category model.Category
	err := c.Db.Where("category_id = ?", categoryID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryServiceImpl) Update(category *model.Category) (*model.Category, error) {
	err := c.Db.Save(category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (c *CategoryServiceImpl) Delete(categoryID string) error {
	var category model.Category
	err := c.Db.Where("category_id = ?", categoryID).Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
