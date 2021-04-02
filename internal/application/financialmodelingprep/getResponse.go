package financialmodelingprep

import (
	"apiStock/internal/structure"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// getResponse Json response from API
func getResponse(listOfCompanies string, arguments structure.Arguments) []byte {
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
		log.Fatal("getResponse function: ", string(responseData))
	}

	return responseData
}

// getResponse Json response from API
func getResponseStockScreener(domain, endpoint, sector, apikey string) []byte {
	url := domain + endpoint + "?sector=" + sector + "&apikey=" + apikey
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

func getResponseStockScreenerWithFilters(arguments structure.Arguments) []byte {
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
