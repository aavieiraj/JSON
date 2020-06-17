package stone

type Items struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	UnitPrice string `json:"unit_price"`
	Quantity  string `json:"quantity"`
	Tangible  bool   `json:"tangible"`
}
