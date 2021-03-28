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
		_ = DiscountedCashFlow(symbol, apikey)

	}
}
