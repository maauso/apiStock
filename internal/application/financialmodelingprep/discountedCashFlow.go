package financialmodelingprep

import (
	"apiStock/internal/arguments"
	"apiStock/internal/domain/persistence"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// discountCashFlow Json arguments
type discountCashFlow struct {
	Symbol     string      `json:"symbol"`
	Date       string      `json:"date"`
	Dcf        interface{} `json:"dcf"`
	StockPrice float64     `json:"Stock Price"`
}

type discountCashFlows []discountCashFlow

// percentageChanged - calculate the percent increase/decrease from two numbers.
// ex. 60.0 is a 200.0% increase from 20.0
func (d *discountCashFlow) percentageChanged() float64 {
	newDcf, err := d.Dcf.(float64)
	if !err {
		fmt.Printf("This values is a String")
		return 0
	}
	return 100 * ((newDcf - d.StockPrice) / d.StockPrice)
}

// DiscountedCashFlowRecover: Recover from redis or financialmodelinggrep the Dcf value
func discountedCashFlowRecover(listOfCompanies string, arguments arguments.Arguments, repo persistence.Repository) discountCashFlows {

	var sb strings.Builder
	var dfc discountCashFlows
	lof := strings.Split(listOfCompanies, ",")

	for _, v := range lof {
		DiscountedCashFlowValue, err := repo.GetTotalCompanies(v)
		if err == nil {
			fmt.Printf("El valor para %v es %v \n", v, DiscountedCashFlowValue)
		} else {
			sb.WriteString(v + ",")
		}
	}

	newList := sb.String()
	if len(newList) > 0 {

		dfc = getDiscountedCashFlowValues(newList, arguments)

		for _, value := range dfc {
			fmt.Printf("Company: %s, Value: %v, DiscountedCashFlowValue: %s \n", value.Symbol, value.StockPrice, value.Dcf)
		}
		populator(dfc, repo)
	}
	return dfc
}

func getDiscountedCashFlowValues(newList string, arguments arguments.Arguments) discountCashFlows {
	var dfc discountCashFlows
	fmt.Printf("Number of companies that we are checking:  %v\n", len(newList))
	responseData := getDiscountedCashFlow(
		newList,
		arguments)
	err := json.Unmarshal(responseData, &dfc)
	if err != nil {
		log.Panic(err)
	}
	return dfc
}
