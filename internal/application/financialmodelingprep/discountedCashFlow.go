package financialmodelingprep

import (
	"encoding/json"
	"fmt"
)

type discountedCashFlow struct {
	Symbol     string  `json:"symbol"`
	Date       string  `json:"date"`
	Dcf        float64 `json:"dcf"`
	StockPrice float64 `json:"Stock Price"`
}

// DiscountedCashFlow just get DFC info
func DiscountedCashFlow(company, apikey string) {
	responseData := getResponse("https://financialmodelingprep.com/",
		"api/v3/company/discounted-cash-flow/",
		company,
		apikey)
	var dfc discountedCashFlow
	_ = json.Unmarshal(responseData, &dfc)
	fmt.Printf("Symbol: %s \nDate: %s\n discountedCashFlow: %v\n StockPrice: %v\n", dfc.Symbol, dfc.Date, dfc.Dcf, dfc.StockPrice)
}
