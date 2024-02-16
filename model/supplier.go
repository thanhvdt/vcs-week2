package model

type Supplier struct {
	//generate fields according to northwind database using gorm
	SupplierID   uint   `json:"supplier_id" gorm:"column:supplier_id;primary_key" validate:"required"`
	CompanyName  string `json:"company_name" gorm:"column:company_name;not null;type:varchar(40);default:null" validate:"required"`
	ContactName  string `json:"contact_name" gorm:"column:contact_name;not null;type:varchar(30);default:null"`
	ContactTitle string `json:"contact_title" gorm:"column:contact_title;not null;type:varchar(30);default:null"`
	Address      string `json:"address" gorm:"column:address;not null;type:varchar(60);default:null"`
	City         string `json:"city" gorm:"column:city;not null;type:varchar(15);default:null"`
	Region       string `json:"region" gorm:"column:region;not null;type:varchar(15);default:null"`
	PostalCode   string `json:"postal_code" gorm:"column:postal_code;not null;type:varchar(10);default:null"`
	Country      string `json:"country" gorm:"column:country;not null;type:varchar(15);default:null"`
	Phone        string `json:"phone" gorm:"column:phone;not null;type:varchar(24);default:null"`
	Fax          string `json:"fax" gorm:"column:fax;not null;type:varchar(24);default:null"`
	HomePage     string `json:"home_page" gorm:"column:home_page;not null;type:text;default:null"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
