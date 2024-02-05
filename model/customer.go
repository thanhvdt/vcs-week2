package model

type Customer struct {
	CustomerID   string `gorm:"column:customer_id;primary_key"`
	CompanyName  string `gorm:"column:company_name"`
	ContactName  string `gorm:"column:contact_name"`
	ContactTitle string `gorm:"column:contact_title"`
	Address      string `gorm:"column:address"`
	City         string `gorm:"column:city"`
	Region       string `gorm:"column:region"`
	PostalCode   string `gorm:"column:postal_code"`
	Country      string `gorm:"column:country"`
	Phone        string `gorm:"column:phone"`
	Fax          string `gorm:"column:fax"`
}

func (Customer) TableName() string {
	return "customers"
}
