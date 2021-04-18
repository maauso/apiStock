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

type DiscountCashFlows []DiscountCashFlow

// PercentageChanged - calculate the percent increase/decrease from two numbers.
// ex. 60.0 is a 200.0% increase from 20.0

// La idea es que la iteraciones con la estructura se realicen a través de metodos, en este caso PercentageChanged se encarga de interactuar
//con los datos de DiscountCashFlow por lo que tiene sentido que 1: esté en el mismo sitio que está la estructura y tenga un metodo tipo estructura
func (d *DiscountCashFlow) PercentageChanged() float64 {
	newDcf, err := d.Dcf.(float64)
	if !err {
		fmt.Printf("This values is a String")
		return 0
	}
	return 100 * ((newDcf - d.StockPrice) / d.StockPrice)
}

// DiscountedCashFlowRetriever just get DFC info
func DiscountedCashFlowRetriever(listOfCompanies string, arguments structure.Arguments, repo persistence.Repository) DiscountCashFlows {

	//var sb strings.Builder
	var dfc DiscountCashFlows
	//lof := strings.Split(listOfCompanies, ",")
	//for _, v := range lof {
	//	if repo.CompanyExists(v) {
	//		fmt.Printf("El valor %s está en Redis\n\n", v)
	//		DiscountedCashFlowValue := repo.GetTotalCompanies(v)
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
		Populator(dfc, repo)
	}
	return dfc
}

func getDiscountedCashFlowValues(newList string, arguments structure.Arguments) DiscountCashFlows {
	var dfc DiscountCashFlows
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
