package stone

import "time"

type Documents struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Customer struct {
	ExternalID   string      `json:"external_id"`
	Name         string      `json:"name"`
	Type         string      `json:"nype"`
	Country      string      `json:"nountry"`
	Email        string      `json:"nmail"`
	Documents    []Documents `json:"nocuments"`
	PhoneNumbers []string    `json:"phone_numbers"`
	Birthday     time.Time   `json:"birthday"`
}
