package category

import "github.com/thanhvdt/vcs-week2/model"

type CategoryRepository interface {
	Save(category *model.Category) (*model.Category, error)
	FindAll() ([]model.Category, error)
	FindByID(categoryID string) (*model.Category, error)
	Update(category *model.Category) (*model.Category, error)
	Delete(categoryID string) error
}
