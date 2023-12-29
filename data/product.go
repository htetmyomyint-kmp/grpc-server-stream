package data

import "math/rand"

type Product struct {
	ID    string
	Price float64
	Name  string
}

type ProductClient struct{}

func NewProductClient() *ProductClient {
	return &ProductClient{}
}

func (p *ProductClient) IsPriceChanged(string) bool {
	return rand.Intn(10)%2 == 0
}
