package stone

import "time"

type Shipping struct {
	Name         string    `json:"name"`
	Fee          string    `json:"fee"`
	DeliveryDate time.Time `json:"delivery_date"`
	Expedited    bool      `json:"expedited"`
	Address      Address   `json:"address"`
}
