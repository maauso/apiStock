package financialmodelingprep

import (
	structs "apiStock/model"
)

// Fmg get metric apikey and company and get the request to financialmodelingprep
func Fmg(p structs.Arguments) {
	switch p.Metric {
	case "DiscountedCashFlow":
		DiscountedCashFlow(p.Company, p.APIKey)
	case "HistoricalDiscountedCashFlow":
		HistoricalDiscountedCashFlow(p.Company, p.APIKey)
	case "KeyMetrics":
		KeyMetrics(p.Company, p.APIKey)
	}
}
