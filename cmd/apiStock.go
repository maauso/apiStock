package main

import (
	"apiStock/pkg/financialmodelingprep/dfc"
	keymetrics "apiStock/pkg/financialmodelingprep/keyMetrics"
)

func main() {
	var apikey = "demo"
	dfc.DiscountedCashFlow(apikey)
	dfc.HistoricalDiscountedCashFlow(apikey)
	keymetrics.KeyMetrics(apikey)
}
