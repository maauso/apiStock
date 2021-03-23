package financialmodelingprep

import (
	"apiStock/internal/domain/laststring"
	"apiStock/internal/domain/persistence"
	"fmt"
	"strings"
)

//UnderValuatedCompanies using dfc method and persistent storage
func UnderValuatedCompanies(apikey string, repo persistence.Repository) {
	symbols := repo.GetData()
	fmt.Println(len(symbols))

	for _, key := range symbols {
		symbol := laststring.LastString(strings.Split(key, " "))
		dfc := DiscountedCashFlow(symbol, apikey)
		fmt.Println(dfc.Symbol)
		fmt.Println(dfc.StockPrice)
		fmt.Println(dfc.Dcf)
	}
}
