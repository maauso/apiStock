package persistence

import (
	"apiStock/internal/structure"
)

type Repository interface {
	// CreateApiStock
	SaveData(dfc structure.DiscountCashFlow)
}
