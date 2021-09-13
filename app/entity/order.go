package entity

type Order struct {
	Milk  float64 `json:"milk,omitempty"`
	Skins int32   `json:"skins,omitempty"`
}

type OrderInput struct {
	Customer string `json:"customer"`
	Order    Order  `json:"order"`
}
