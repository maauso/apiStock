package persistence

import (
	"github.com/zenthangplus/goccm"
)

//Repository Interface repository for a persistence
type Repository interface {

	// El repository esta ligado al objeto
	PopulateData(Symbol string, DiscountCashFlow interface{}, c goccm.ConcurrencyManager)
	PopulateUnderValuatedCompanies(Symbol string, Percentage float64, c goccm.ConcurrencyManager)
	GetValue(listOfElements string) float64
}
