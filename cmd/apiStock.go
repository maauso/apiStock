package main

import (
	structs "apiStock/model"
	"apiStock/pkg/selector"
	"apiStock/tools"
	"flag"
	"os"
)

func main() {

	providerPtr := tools.DerefString(flag.String("provider", "financialmodelingprep", "Select the provider [ financialmodelingprep ] (Required)"))
	metricPtr := tools.DerefString(flag.String("metric", "DiscountedCashFlow", "Metric { DiscountedCashFlow | HistoricalDiscountedCashFlow | KeyMetric };(Required)"))
	companySymbolPtr := tools.DerefString(flag.String("company", "AAPL", "demo apikey"))
	apikeyPtr := tools.DerefString(flag.String("apikey", "demo", "demo apikey"))

	flag.Parse()

	if providerPtr == "" || metricPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	params := structs.Arguments{Provider: providerPtr, Metric: metricPtr, Company: companySymbolPtr, APIKey: apikeyPtr}
	selector.GetMetric(params)
}
