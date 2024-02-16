package category

import "github.com/thanhvdt/vcs-week2/model"

type CategoryService interface {
	Create(category *model.Category) (*model.Category, error)
	ReadAll() ([]model.Category, error)
	ReadByID(categoryID string) (*model.Category, error)
	Update(category *model.Category) (*model.Category, error)
	Delete(categoryID string) error
}
