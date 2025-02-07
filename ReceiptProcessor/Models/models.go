package models

type Receipt struct {
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Total        string `json:"total" validate:"required"`
	Items        []Item `json:"items" validate:"required,min=1"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required"`
}

func (r *Receipt) SetRetailer(newRetailer string) {
	r.Retailer = newRetailer
}
