package supplier

import (
	"github.com/go-playground/validator/v10"
	"github.com/thanhvdt/vcs-week2/model"
	"github.com/thanhvdt/vcs-week2/repository/supplier"
)

type SupplierServiceImpl struct {
	supplierRepository supplier.SupplierRepository
	Validate           *validator.Validate
}

func NewSupplierService(supplierRepository supplier.SupplierRepository, validate *validator.Validate) *SupplierServiceImpl {
	return &SupplierServiceImpl{supplierRepository: supplierRepository, Validate: validate}
}

func (s *SupplierServiceImpl) Create(supplier *model.Supplier) (*model.Supplier, error) {
	err := s.Validate.Struct(supplier)
	if err != nil {
		return nil, err
	}
	return s.supplierRepository.Save(supplier)
}

func (s *SupplierServiceImpl) ReadAll() ([]model.Supplier, error) {
	return s.supplierRepository.FindAll()
}

func (s *SupplierServiceImpl) ReadByID(supplierID string) (*model.Supplier, error) {
	return s.supplierRepository.FindByID(supplierID)
}

func (s *SupplierServiceImpl) Update(supplierID string, supplier *model.Supplier) (*model.Supplier, error) {
	if err := s.Validate.Struct(supplier); err != nil {
		return nil, err
	}
	return s.supplierRepository.Update(supplier)
}

func (s *SupplierServiceImpl) Delete(supplierID string) error {
	return s.supplierRepository.Delete(supplierID)
}
