package models

type Product struct {
	Id          int64
	Name        string
	Description string
	Stock       int64
	Price       float64
	Provider    Provider
}
