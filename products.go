package go_project

type Products struct {
  Id int `json:"id"`
  Producttitle string `json:"producttitle"`
  Price int `json:"price"`
  Quantity int `json:"quantity"`
  Pddesc string `json:"pddesc"`
  Image string `json:"image"`
}