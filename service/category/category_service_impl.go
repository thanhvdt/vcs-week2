package category

import (
	"github.com/go-playground/validator/v10"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/repository/category"
)

type CategoryServiceImpl struct {
	CategoryRepository category.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository category.CategoryRepository, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, Validate: validate}
}

func (c *CategoryServiceImpl) Create(category *model.Category) (*model.Category, error) {
	err := c.Validate.Struct(category)
	if err != nil {
		return nil, err
	}
	return c.CategoryRepository.Save(category)
}

func (c *CategoryServiceImpl) ReadAll() ([]model.Category, error) {
	return c.CategoryRepository.FindAll()
}

func (c *CategoryServiceImpl) ReadByID(categoryID string) (*model.Category, error) {
	return c.CategoryRepository.FindByID(categoryID)
}

func (c *CategoryServiceImpl) Update(category *model.Category) (*model.Category, error) {
	if err := c.Validate.Struct(category); err != nil {
		return nil, err
	}
	return c.CategoryRepository.Update(category)
}

func (c *CategoryServiceImpl) Delete(categoryID string) error {
	return c.CategoryRepository.Delete(categoryID)
}
