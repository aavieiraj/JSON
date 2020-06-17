package stone

type Address struct {
	Country      string `json:"country"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
	Zipcode      string `json:"zipcode"`
}
