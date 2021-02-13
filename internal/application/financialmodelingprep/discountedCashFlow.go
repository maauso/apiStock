package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"encoding/json"
	"fmt"
)

// DiscountedCashFlow just get DFC info
func DiscountedCashFlow(company, apikey string, repo persistence.Repository) {
	responseData := getResponse("https://financialmodelingprep.com/",
		"api/v3/company/discounted-cash-flow/",
		company,
		apikey)
	var dfc structure.DiscountCashFlow
	_ = json.Unmarshal(responseData, &dfc)

	fmt.Printf("Symbol: %s \nDate: %s\n DiscountCashFlow: %v\n StockPrice: %v\n", dfc.Symbol, dfc.Date, dfc.Dcf, dfc.StockPrice)
	repo.SaveData(dfc)
}
