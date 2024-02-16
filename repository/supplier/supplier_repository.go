package supplier

import "github.com/thanhvdt/vcs-week2/model"

type SupplierRepository interface {
	Save(supplier *model.Supplier) (*model.Supplier, error)
	FindAll() ([]model.Supplier, error)
	FindByID(supplierID string) (*model.Supplier, error)
	Update(supplier *model.Supplier) (*model.Supplier, error)
	Delete(supplierID string) error
}
