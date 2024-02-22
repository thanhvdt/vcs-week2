package customer

import (
	"github.com/thanhvdt/vcs-week2/model"
)

type CustomerRepository interface {
	Save(customer *model.Customer) (*model.Customer, error)
	FindAll() ([]model.Customer, error)
	FindByID(customerID string) (*model.Customer, error)
	Update(customer *model.Customer) (*model.Customer, error)
	Delete(customerID string) error
	SearchByCompany(company string) (map[string]interface{}, error)
}
