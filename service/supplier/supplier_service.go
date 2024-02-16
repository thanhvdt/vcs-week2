package supplier

import "github.com/thanhvdt/vcs-week2/model"

type SupplierService interface {
	Create(supplier *model.Supplier) (*model.Supplier, error)
	ReadAll() ([]model.Supplier, error)
	ReadByID(supplierID string) (*model.Supplier, error)
	Update(supplierID string, supplier *model.Supplier) (*model.Supplier, error)
	Delete(supplierID string) error
}
