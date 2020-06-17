package stone

type Transaction struct {
	Amount             string   `json:"amount"`
	CardNumber         string   `json:"card_number"`
	CardCVV            string   `json:"card_cvv"`
	CardExpirationDate string   `json:"card_expiration_date"`
	CardHolderName     string   `json:"card_holder_name"`
	Customer           Customer `json:"customer"` // Custumer da transação
	Biiling            Billing  `json:"billing"`  // Endereço de cobrança
	Shipping           Shipping `json:"shipping"` // Endereço de envio
	Items              []Items  `json:"items"`    // Vetor com itens da transação
}
