package models

type Product struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating float32 `json:"rating"`
}

func MockProducts() (products [4]Product) {
	products = [4]Product{
		CProduct("iPhone8", 45_000, 4.5),
		CProduct("iPhone9", 54_000, 4.5),
		CProduct("iPhoneX", 58_000, 4.5),
		CProduct("iPhone11", 85_000, 4.5),
	}
	return
}

func CProduct(name string, price float64, rating float32) Product {
	return Product{name, price, rating}
}
