package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type companiesFiltered []struct {
	Symbol             string  `json:"symbol"`
	CompanyName        string  `json:"companyName"`
	MarketCap          float64 `json:"marketCap"`
	Sector             string  `json:"sector"`
	Industry           string  `json:"industry"`
	Beta               float64 `json:"beta"`
	Price              float64 `json:"price"`
	LastAnnualDividend float64 `json:"lastAnnualDividend"`
	Volume             int     `json:"volume"`
	Exchange           string  `json:"exchange"`
	ExchangeShortName  string  `json:"exchangeShortName"`
	Country            string  `json:"country"`
	IsEtf              bool    `json:"isEtf"`
	IsActivelyTrading  bool    `json:"isActivelyTrading"`
}

//UnderValuatedCompanies using dfc method and persistent storage
func UnderValuatedCompanies(arguments structure.Arguments, repo persistence.Repository) {
	var cf companiesFiltered
	var sb strings.Builder
	responseData := getResponseStockScreenerWithFilters(arguments)
	err := json.Unmarshal(responseData, &cf)
	if err != nil {
		log.Panic(err)
	}
	for _, value := range cf {
		sb.WriteString(value.Symbol + ",")
	}
	listOfCompanies := sb.String()
	listOfCompanies = strings.TrimSuffix(listOfCompanies, ",")

	dfc := DiscountedCashFlow(listOfCompanies, arguments, repo)
	for _, value := range dfc {
		growth := PercentageChanged(value.StockPrice, value.Dcf)
		if growth >= *arguments.PercentageOfGrowth {
			fmt.Printf(
				"Symbol: %v, StockPrice: %v , DiscountCashFlowValue: %v,  Change : %0.2f %% \n\n",
				value.Symbol,
				value.StockPrice,
				value.Dcf,
				growth,
			)
		}
	}
}

// PercentageChanged - calculate the percent increase/decrease from two numbers.
// ex. 60.0 is a 200.0% increase from 20.0
func PercentageChanged(StockPrice float64, Dcf interface{}) float64 {
	newDcf, err := Dcf.(float64)
	if !err {
		fmt.Printf("This values is a String")
		return 0
	}
	return 100 * ((newDcf - StockPrice) / StockPrice)
}
