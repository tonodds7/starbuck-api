package models

type Item struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type ItemDetail struct {
	ID                    int
	Name                  string
	Detail                string
	Category              string
	Acidity               int
	Body                  int
	Processing_Method     string
	Tasting_Notes         string
	Complementary_Flavors string
	Region                string
	Price                 float64
}

type RequestOrder struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}
