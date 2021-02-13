package persistence

import (
	"apiStock/internal/structure"
)

//Repository Interface repository for a persistence
type Repository interface {
	// CreateApiStock
	SaveData(dfc structure.DiscountCashFlow)
}
