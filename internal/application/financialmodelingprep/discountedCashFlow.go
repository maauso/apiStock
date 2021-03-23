package financialmodelingprep

import (
	"apiStock/internal/structure"
	"encoding/json"
)

// DiscountedCashFlow just get DFC info
func DiscountedCashFlow(company, apikey string) structure.DiscountCashFlow {
	responseData := getResponse("https://financialmodelingprep.com/",
		"api/v3/company/discounted-cash-flow/",
		company,
		apikey)
	var dfc structure.DiscountCashFlow
	_ = json.Unmarshal(responseData, &dfc)
	return dfc
}
