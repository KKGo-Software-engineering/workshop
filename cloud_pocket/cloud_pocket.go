package cloud_pocket

import (
	"math/big"
	"time"
)

type Currency string

const (
	THBCurrency Currency = "THB"
)

type CloudPocket struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Amount    big.Float `json:"amount"`
	Goal      big.Float `json:"goal"`
	Currency  Currency  `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}
