package repository

import "github.com/thanhvdt/vcs-week2/model"

type CategoryRepository interface {
	Save(category *model.Categories) (*model.Categories, error)
	FindAll() ([]model.Categories, error)
	FindByID(categoryID string) (*model.Categories, error)
	Update(category *model.Categories) (*model.Categories, error)
	Delete(categoryID string) error
}
