package persistence

import (
	"github.com/zenthangplus/goccm"
)

//Repository Interface repository for a persistence
type Repository interface {
	PopulateData(Symbol string, DiscountCashFlow interface{}, c goccm.ConcurrencyManager)

	CompanyExists(listOfElements string) bool
	PopulateUnderValuatedCompanies(Symbol string, Percentage float64, c goccm.ConcurrencyManager)
	GetTotalCompanies(listOfCompanyNames string) float64
}
