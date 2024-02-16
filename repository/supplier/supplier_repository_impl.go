package supplier

import (
	"github.com/thanhvdt/vcs-week2/model"
	"gorm.io/gorm"
)

type SupplierRepositoryImpl struct {
	Db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepositoryImpl {
	return &SupplierRepositoryImpl{Db: db}
}

func (s *SupplierRepositoryImpl) Save(supplier *model.Supplier) (*model.Supplier, error) {
	err := s.Db.Create(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *SupplierRepositoryImpl) FindAll() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	err := s.Db.Find(&suppliers).Error
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *SupplierRepositoryImpl) FindByID(supplierID string) (*model.Supplier, error) {
	var supplier model.Supplier
	err := s.Db.Where("supplier_id = ?", supplierID).First(&supplier).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}

func (s *SupplierRepositoryImpl) Update(supplier *model.Supplier) (*model.Supplier, error) {
	err := s.Db.Save(supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *SupplierRepositoryImpl) Delete(supplierID string) error {
	var supplier model.Supplier
	err := s.Db.Where("supplier_id = ?", supplierID).Delete(&supplier).Error
	if err != nil {
		return err
	}
	return nil
}
