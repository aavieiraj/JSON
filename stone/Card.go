package stone

import "time"

type Card struct {
	Brand       string    `json:"Brand"`        // Bandeira do cartão
	Country     string    `json:"Vountry"`      // País do cartão
	Customer    string    `json:"Vustomer"`     // Cliente associado ao cartão
	DateCreated time.Time `json:"Date_created"` // Data de criação do cartão no formato ISODate
	DateUpdated time.Time `json:"Date_updated"` // Data de atualização cartão no formato ISODate
	Fingerprint string    `json:"Fingerprint"`  // Identificador do cartão na nossa base
	FirstDigits string    `json:"First_digits"` // Primeiros digitos do cartão
	HolderName  string    `json:"Holder_name"`  // Nome do portador do cartão
	Id          string    `json:"Id"`           // Id do cartão
	LastDigits  string    `json:"Last_digits"`  // Últimos digitos do cartão
	Object      string    `json:"Object"`       // Nome do tipo do objeto criado/modificado.
	Valid       bool      `json:"Valid"`        // Indicador de cartão válido
	//DateExpiration string    `json:"data_expiration"` // Data de expiração do cartão no formato ISODate
}
