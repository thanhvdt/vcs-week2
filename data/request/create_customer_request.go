package request

type CreateCustomerRequest struct {
	CustomerID   string `json:"customer_id" validate:"required"`
	CompanyName  string `json:"company_name" validate:"required"`
	ContactName  string `json:"contact_name"`
	ContactTitle string `json:"contact_title"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Region       string `json:"region"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
}
