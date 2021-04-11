package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
)

// Fmg get metric apikey and company and get the request to financialmodelingprep
func Fmg(arguments *structure.Arguments, repo persistence.Repository) {
	switch *arguments.Metric {
	case "DiscountedCashFlow":
		DiscountedCashFlow(*arguments.Company, *arguments, repo)
	case "UnderValuatedCompanies":
		UnderValuatedCompanies(*arguments, repo)
	}
}
