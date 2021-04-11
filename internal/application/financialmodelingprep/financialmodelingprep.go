package financialmodelingprep

import (
	"apiStock/internal/arguments"
	"apiStock/internal/domain/persistence"
)

// Fmg get metric apikey and company and get the request to financialmodelingprep
func Fmg(arguments *arguments.Arguments, repo persistence.Repository) {
	switch *arguments.Metric {
	case "DiscountedCashFlow":
		discountedCashFlowRecover(*arguments.Company, *arguments, repo)
	case "UnderValuatedCompanies":
		underValuatedCompanies(*arguments, repo)
	}
}
