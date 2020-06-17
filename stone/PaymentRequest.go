package stone

//PaymentRequest is a request to process
type PaymentRequest struct {
	Action      string      `json:"Action"`
	Transaction Transaction `json:"Sale"`
}
