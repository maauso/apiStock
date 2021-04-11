package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"

	"github.com/zenthangplus/goccm"
)

//type stockScreener []struct {
//	Symbol             string  `json:"symbol"`
//	CompanyName        string  `json:"companyName"`
//	MarketCap          int64   `json:"marketCap"`
//	Sector             string  `json:"sector"`
//	Industry           string  `json:"industry"`
//	Beta               float64 `json:"beta"`
//	Price              float64 `json:"price"`
//	LastAnnualDividend int     `json:"lastAnnualDividend"`
//	Volume             int     `json:"volume"`
//	Exchange           string  `json:"exchange"`
//	ExchangeShortName  string  `json:"exchangeShortName"`
//	Country            string  `json:"country"`
//	IsEtf              bool    `json:"isEtf"`
//	IsActivelyTrading  bool    `json:"isActivelyTrading"`
//}

//Populator get a group of companies
func Populator(dfc DiscountCashFlow, repo persistence.Repository) {
	c := goccm.New(10)
	for _, value := range dfc {
		c.Wait()
		go repo.PopulateData(value.Symbol, value.Dcf, c)
	}
	c.WaitAllDone()
}
