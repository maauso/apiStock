package financialmodelingprep

import (
	"apiStock/internal/structure"
	"encoding/json"
	"fmt"
	"log"
)

// DiscountCashFlow Json structure
type DiscountCashFlow []struct {
	Symbol     string      `json:"symbol"`
	Date       string      `json:"date"`
	Dcf        interface{} `json:"dcf"`
	StockPrice float64     `json:"Stock Price"`
}

// DiscountedCashFlow just get DFC info
func DiscountedCashFlow(listOfCompanies string, arguments structure.Arguments) DiscountCashFlow {
	var dfc DiscountCashFlow

	fmt.Printf("Number of companies that we are checking:  %v\n", len(listOfCompanies))
	responseData := getResponse(
		listOfCompanies,
		arguments)
	err := json.Unmarshal(responseData, &dfc)
	if err != nil {
		log.Panic(err)
	}
	if *arguments.Metric == "DiscountedCashFlow" {
		for _, value := range dfc {
			fmt.Printf(
				"Symbol: %v, StockPrice: %v , DiscountCashFlowValue: %v,",
				value.Symbol,
				value.StockPrice,
				value.Dcf,
			)
		}
	}
	return dfc
}
