package persistence

import (
	"github.com/zenthangplus/goccm"
)

//Repository Interface repository for a persistence
type Repository interface {
	// CreateApiStock
	GetData() []string
	PopulateData(Sector, Symbol, CompanyName string, c goccm.ConcurrencyManager)
	PopulateUnderValuatedCompanies(Symbol string, Percentage float64, c goccm.ConcurrencyManager)
}
