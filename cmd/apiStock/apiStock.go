package main

import (
	"apiStock/internal/application/selector"
	"apiStock/internal/domain/derefstring"
	"apiStock/pkg/infrastructure/http/arguments"
	"flag"
	"os"
)

func main() {

	providerPtr := derefstring.DerefString(flag.String("provider", "financialmodelingprep", "Select the provider [ financialmodelingprep ] (Required)"))
	metricPtr := derefstring.DerefString(flag.String("metric", "DiscountedCashFlow", "Metric { DiscountedCashFlow | HistoricalDiscountedCashFlow | KeyMetric };(Required)"))
	companySymbolPtr := derefstring.DerefString(flag.String("company", "AAPL", "demo apikey"))
	apikeyPtr := derefstring.DerefString(flag.String("apikey", "demo", "demo apikey"))

	flag.Parse()

	if providerPtr == "" || metricPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	params := arguments.Arguments{Provider: providerPtr, Metric: metricPtr, Company: companySymbolPtr, APIKey: apikeyPtr}
	selector.GetMetric(params)
}
