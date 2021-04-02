package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"encoding/json"

	"github.com/zenthangplus/goccm"
)

type stockScreener []struct {
	Symbol             string  `json:"symbol"`
	CompanyName        string  `json:"companyName"`
	MarketCap          int64   `json:"marketCap"`
	Sector             string  `json:"sector"`
	Industry           string  `json:"industry"`
	Beta               float64 `json:"beta"`
	Price              float64 `json:"price"`
	LastAnnualDividend int     `json:"lastAnnualDividend"`
	Volume             int     `json:"volume"`
	Exchange           string  `json:"exchange"`
	ExchangeShortName  string  `json:"exchangeShortName"`
	Country            string  `json:"country"`
	IsEtf              bool    `json:"isEtf"`
	IsActivelyTrading  bool    `json:"isActivelyTrading"`
}

//PopulateCompanies get a group of companies
func PopulateCompanies(apikey string, repo persistence.Repository) {
	c := goccm.New(10)
	responseData := getResponseStockScreener("https://financialmodelingprep.com/api/v3/", "stock-screener", "Technology", apikey)
	var screener stockScreener
	_ = json.Unmarshal(responseData, &screener)
	for _, value := range screener {
		c.Wait()
		go repo.PopulateData(value.Sector, value.Symbol, value.CompanyName, c)
	}
	c.WaitAllDone()
}
