package models

type Region struct {
	Name int
}

type City struct {
	ID         int    `json:"id"`         //(уникальный номер)
	Name       string `json:"name"`       //(название города)
	Region     string `json:"region"`     //(регион)
	District   string `json:"district"`   //(округ)
	Population int    `json:"population"` //(численность населения)
	Foundation int    `json:"foundation"` //(год основания)
}
