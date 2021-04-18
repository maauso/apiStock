package persistence

import (
	"github.com/zenthangplus/goccm"
)

//Repository Interface repository for a persistence
type Repository interface {

	// El repository esta ligado al objeto
	PopulateData(Symbol string, DiscountCashFlow interface{}, c goccm.ConcurrencyManager)

	// CreateApiStock
	CompanyExists(listOfElements string) bool
	PopulateUnderValuatedCompanies(Symbol string, Percentage float64, c goccm.ConcurrencyManager)
	GetTotalCompanies(listOfCompanyNames string) float64

	//este string tiene que ser un [] si a la api lo hace falta una lista la creas en aplicacion.

	// La idea es que el repositorio esté reflejada la lógica de la aplicacion, indistintamente de la base de datos
}
