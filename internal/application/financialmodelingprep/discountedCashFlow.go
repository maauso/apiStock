package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"encoding/json"
	"fmt"
	"log"
)

// DiscountCashFlow Json structure
type DiscountCashFlow struct {
	Symbol     string      `json:"symbol"`
	Date       string      `json:"date"`
	Dcf        interface{} `json:"dcf"`
	StockPrice float64     `json:"Stock Price"`
}

// DiscountedCashFlow just get DFC info
func DiscountedCashFlow(listOfCompanies string, arguments structure.Arguments, repo persistence.Repository) DiscountCashFlow {

	for _, v := range *d {
		if v.StockPrice > 10 {
			top3 = append(top3, v)
		}
	}
	return top3
}

// DiscountedCashFlowRetriever just get DFC info
func DiscountedCashFlowRetriever(listOfCompanies string, arguments structure.Arguments, repo persistence.Repository) DiscountCashFlows {

	//var sb strings.Builder
	var dfc DiscountCashFlows
	//lof := strings.Split(listOfCompanies, ",")
	//for _, v := range lof {
	//	if repo.ComanyExists(v) {
	//		fmt.Printf("El valor %s está en Redis\n\n", v)
	//		DiscountedCashFlowValue := repo.GetTotalComanies(v)
	//		fmt.Printf("El valor para %v es %v ", v, DiscountedCashFlowValue)
	//	} else {
	//		sb.WriteString(v + ",")
	//	}
	//}
	//newList := sb.String()
	if len(listOfCompanies) > 0 {

		dfc = getDiscountedCashFlowValues(listOfCompanies, arguments)

		for _, value := range dfc {
			fmt.Printf("Company: %s, Value: %v, DiscountedCashFlowValue: %s", value.Symbol, value.StockPrice, value.Dcf)
		}
		repo.Populator(dfc)
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
