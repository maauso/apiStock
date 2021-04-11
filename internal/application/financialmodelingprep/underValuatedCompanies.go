package financialmodelingprep

import (
	"apiStock/internal/arguments"
	"apiStock/internal/domain/persistence"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type companyFiltered struct {
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

type companiesFiltered []companyFiltered

//underValuatedCompanies using dfc method and persistent storage
func underValuatedCompanies(arguments arguments.Arguments, repo persistence.Repository) {
	var cf companiesFiltered
	var sb strings.Builder
	responseData := getCompaniesStockScreener(arguments)
	err := json.Unmarshal(responseData, &cf)
	if err != nil {
		log.Panic(err)
	}
	for _, value := range cf {
		sb.WriteString(value.Symbol + ",")
	}
	listOfCompanies := sb.String()
	listOfCompanies = strings.TrimSuffix(listOfCompanies, ",")

	groupOfDiscountedFlow := discountedCashFlowRecover(listOfCompanies, arguments, repo)
	for _, dfc := range groupOfDiscountedFlow {
		growth := dfc.PercentageChanged()
		if growth >= *arguments.PercentageOfGrowth {
			fmt.Printf(
				"Symbol: %v, StockPrice: %v , DiscountCashFlowValue: %v,  Change : %0.2f %% \n\n",
				dfc.Symbol,
				dfc.StockPrice,
				dfc.Dcf,
				growth,
			)
		}
	}
}
