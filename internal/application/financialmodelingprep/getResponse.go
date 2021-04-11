package financialmodelingprep

import (
	"apiStock/internal/arguments"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// getDiscountedCashFlow value using financialmodelingprep API
func getDiscountedCashFlow(listOfCompanies string, arguments arguments.Arguments) []byte {
	var (
		url = "https://" + *arguments.Provider + ".com/api/v3/" + "discounted-cash-flow/" +
			listOfCompanies +
			"?apikey=" + *arguments.APIKey
	)
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
	}
	if strings.Contains(string(responseData), "Error Message") {
		log.Fatal("getDiscountedCashFlow function: ", string(responseData))
	}

	return responseData
}

//getCompaniesStockScreener get list of companies, filtering per sector/beta/marketCap through financialmodelingprep
func getCompaniesStockScreener(arguments arguments.Arguments) []byte {
	var (
		url = "https://" + *arguments.Provider + ".com/api/v3/" + "stock-screener" +
			"?marketCapMoreThan=" + *arguments.UnderValuedArguments.MarketCap +
			"&sector=" + *arguments.UnderValuedArguments.Sector +
			"&beta=" + *arguments.UnderValuedArguments.Beta +
			"&apikey=" + *arguments.APIKey
	)
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
	}
	return responseData
}
