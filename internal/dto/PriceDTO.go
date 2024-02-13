package dto

type PriceDTO struct {
	ID       string      `json:"id"`
	Amount   float64     `json:"amount"`
	Currency CurrencyDTO `json:"currency"`
}

type CurrencyDTO struct {
	ID               string `json:"id"`
	Symbol           string `json:"symbol"`
	DecimalDivider   string `json:"decimal_divider"`
	ThousandsDivider string `json:"thousands_divider"`
}
