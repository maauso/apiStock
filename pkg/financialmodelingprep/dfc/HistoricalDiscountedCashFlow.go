package dfc

import (
	getjsonresponse "apiStock/pkg/getJSONResponse"
	"encoding/json"
	"fmt"
)

type historicalDiscountedCashFlow []struct {
	Symbol        string `json:"symbol"`
	HistoricalDCF []struct {
		Date       string  `json:"date"`
		StockPrice float64 `json:"Stock Price"`
		DCF        float64 `json:"DCF"`
	} `json:"historicalDCF"`
}

//HistoricalDiscountedCashFlow by stock
func HistoricalDiscountedCashFlow(apikey string) {

	responseData := getjsonresponse.GetResponse("https://financialmodelingprep.com/",
		"api/v3/historical-discounted-cash-flow/",
		"AAPL",
		"demo")

	var historicalDfc historicalDiscountedCashFlow

	json.Unmarshal([]byte(responseData), &historicalDfc)

	for _, value := range historicalDfc {

		fmt.Println(value.Symbol)
		for _, v := range value.HistoricalDCF {

			fmt.Println(v.DCF)
			fmt.Println(v.Date)
			fmt.Println(v.StockPrice)

		}

	}

}
