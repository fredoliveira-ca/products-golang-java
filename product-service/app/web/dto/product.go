package dto

//Product is ...
type Product struct {
	ID           string
	PriceInCents int64
	Title        string
	Description  string
	Discount     Discount
}
