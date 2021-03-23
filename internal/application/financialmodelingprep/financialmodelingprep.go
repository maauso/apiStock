package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
)

// Fmg get metric apikey and company and get the request to financialmodelingprep
func Fmg(p structure.Arguments, repo persistence.Repository) {
	if p.ListOfCompanies == "no" {
		switch p.Metric {
		case "DiscountedCashFlow":
			DiscountedCashFlow(p.Company, p.APIKey)
		case "HistoricalDiscountedCashFlow":
			HistoricalDiscountedCashFlow(p.Company, p.APIKey)
		case "KeyMetrics":
			KeyMetrics(p.Company, p.APIKey)
		case "UnderValuatedCompanies":
			UnderValuatedCompanies(p.APIKey, repo)
		}
	} else {
		PopulateCompanies(p.APIKey, repo)
	}
}
