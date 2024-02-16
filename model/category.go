package model

type Category struct {
	CategoryID   uint   `json:"category_id" gorm:"column:category_id;primary_key" validate:"required"`
	CategoryName string `json:"category_name" gorm:"column:category_name;not null;type:varchar(15);default:null" validate:"required"`
	Description  string `json:"description" gorm:"column:description"`
	Picture      []byte `json:"picture" gorm:"column:picture"`
}
