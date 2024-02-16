package model

type Product struct {
	ProductID       uint    `json:"product_id" gorm:"column:product_id;primary_key" validate:"required"`
	ProductName     string  `json:"product_name" gorm:"column:product_name;not null;type:varchar(40);default:null" validate:"required"`
	SupplierID      *uint   `json:"supplier_id" gorm:"column:supplier_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID      *uint   `json:"category_id" gorm:"column:category_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	QuantityPerUnit string  `json:"quantity_per_unit" gorm:"column:quantity_per_unit;not null;type:varchar(20);default:null"`
	UnitPrice       float64 `json:"unit_price" gorm:"column:unit_price;not null;type:decimal(10,2);default:null"`
	UnitsInStock    uint    `json:"units_in_stock" gorm:"column:units_in_stock;not null;default:null"`
	UnitsOnOrder    uint    `json:"units_on_order" gorm:"column:units_on_order;not null;default:null"`
	ReorderLevel    uint    `json:"reorder_level" gorm:"column:reorder_level;not null;default:null"`
	Discontinued    uint    `json:"discontinued" gorm:"column:discontinued;not null;default:null"`

	Supplier Supplier `gorm:"foreignKey:SupplierID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}

func (Product) TableName() string {
	return "products"
}
