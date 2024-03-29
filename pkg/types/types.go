package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string

// Category ...
type Category string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB	Currency = "RUB"
	USD Currency = "USD"
	UZB Currency = "UZB"
)

// PAN представляет номер карты
type PAN string

// 	Card представляет информацию о платёжной карте
type Card struct{
	ID	int
	PAN PAN
	Balance Money // использовали Money
	MinBalance	Money
	Currency Currency
	Color	string
	Name string
	Active bool
}

// Payment представляет информацию о платеже.
type Payment struct{
	ID int
	Amount Money
	Category Category
}

type PaymentSource struct{
	Type	string 		// card
	Number	string 		// номер вида '5058 xxxx xxxx 8888'
	Balance Money // баланс в дирамах
}