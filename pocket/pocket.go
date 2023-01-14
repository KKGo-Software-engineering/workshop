package pocket

type Pocket struct {
	Id       int     `json:"id"`
	Name     string  `json:"title"`
	Category string  `json:"amount"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}
