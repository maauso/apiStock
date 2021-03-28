package financialmodelingprep

import (
	"apiStock/internal/structure"
	"encoding/json"
	"fmt"
	"log"
)

// DiscountedCashFlow just get DFC info
func DiscountedCashFlow(company, apikey string) structure.DiscountCashFlow {
	var dfc structure.DiscountCashFlow
	responseData := getResponse("https://financialmodelingprep.com/",
		"api/v3/discounted-cash-flow/",
		company,
		apikey)

	err := json.Unmarshal(responseData, &dfc)
	if err != nil {
		log.Panic(err)
	}
	for _, value := range dfc {
		fmt.Printf("Symbol: %v , StockPrice: %v , DiscountCashFlowValue: %v ", value.Symbol, value.StockPrice, value.Dcf)
	}
	return dfc
}
