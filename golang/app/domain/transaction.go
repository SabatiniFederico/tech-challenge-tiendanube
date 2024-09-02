package domain

type Transaction struct {
	Value              string
	Description        string
	Method             string
	CardNumber         string
	CardExpirationDate string
	CardCvv            string
	Idd                string
}
