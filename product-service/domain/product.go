package dto

//Product domain struct
type Product struct {
	ID           string
	PriceInCents int64
	Title        string
	Description  string
	Discount     Discount
}
