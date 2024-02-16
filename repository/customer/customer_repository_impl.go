package customer

import (
	"github.com/thanhvdt/vcs-week2/model"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	Db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{Db: db}
}

func (c *CustomerRepositoryImpl) Save(customer *model.Customer) (*model.Customer, error) {
	err := c.Db.Create(customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerRepositoryImpl) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	err := c.Db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *CustomerRepositoryImpl) FindByID(customerID string) (*model.Customer, error) {
	var customer model.Customer
	err := c.Db.Where("customer_id = ?", customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *CustomerRepositoryImpl) Update(customer *model.Customer) (*model.Customer, error) {
	err := c.Db.Save(customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerRepositoryImpl) Delete(customerID string) error {
	var customer model.Customer
	err := c.Db.Where("customer_id = ?", customerID).Delete(&customer).Error
	if err != nil {
		return err
	}
	return nil
}
