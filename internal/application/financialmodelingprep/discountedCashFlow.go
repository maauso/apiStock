package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// DiscountCashFlow Json structure
type DiscountCashFlow []struct {
	Symbol     string      `json:"symbol"`
	Date       string      `json:"date"`
	Dcf        interface{} `json:"dcf"`
	StockPrice float64     `json:"Stock Price"`
}

// DiscountedCashFlow just get DFC info
func DiscountedCashFlow(listOfCompanies string, arguments structure.Arguments, repo persistence.Repository) DiscountCashFlow {

	var sb strings.Builder
	var dfc DiscountCashFlow
	lof := strings.Split(listOfCompanies, ",")

	for _, v := range lof {
		if repo.KeyExist(v) {
			fmt.Printf("El valor %s está en Redis\n\n", v)
			DiscountedCashFlowValue := repo.GetValue(v)
			fmt.Printf("El valor para %v es %v ", v, DiscountedCashFlowValue)
		} else {
			sb.WriteString(v + ",")
		}
	}
	newList := sb.String()
	if len(newList) > 0 {

		dfc := getDiscountedCashFlowValues(newList, arguments)

		for _, value := range dfc {
			fmt.Printf("Company: %s, Value: %v, DiscountedCashFlowValue: %s", value.Symbol, value.StockPrice, value.Dcf)
		}
		Populator(dfc, repo)
	}
	return dfc
}

func getDiscountedCashFlowValues(newList string, arguments structure.Arguments) DiscountCashFlow {
	var dfc DiscountCashFlow
	fmt.Printf("Number of companies that we are checking:  %v\n", len(newList))
	responseData := getResponse(
		newList,
		arguments)
	err := json.Unmarshal(responseData, &dfc)
	if err != nil {
		log.Panic(err)
	}
	return dfc
}
