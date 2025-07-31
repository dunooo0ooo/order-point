package entity

import "time"

type ProductType string

const (
	ProductTypeElectronics ProductType = "electronics"
	ProductTypeClothing    ProductType = "clothing"
	ProductTypeShoes       ProductType = "shoes"
)

type Product struct {
	Id             int64
	AcceptanceDate time.Time
	Type           ProductType
}
