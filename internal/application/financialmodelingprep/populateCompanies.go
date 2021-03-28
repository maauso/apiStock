package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	json "encoding/json"

	"github.com/zenthangplus/goccm"
)

//PopulateCompanies get a group of companies
func PopulateCompanies(apikey string, repo persistence.Repository) {
	c := goccm.New(10)
	responseData := getResponseStockScreener("https://financialmodelingprep.com/api/v3/", "stock-screener", "Technology", apikey)
	var screener structure.StockScreener
	_ = json.Unmarshal(responseData, &screener)
	for _, value := range screener {
		c.Wait()
		go repo.PopulateData(value.Sector, value.Symbol, value.CompanyName, c)
	}
	c.WaitAllDone()
}
