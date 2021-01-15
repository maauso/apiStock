package keymetrics

import (
	getjsonresponse "apiStock/pkg/getJSONResponse"
	"encoding/json"
	"fmt"
)

type keyMetrics []struct {
	Symbol                                 string      `json:"symbol"`
	Date                                   string      `json:"date"`
	RevenuePerShare                        float64     `json:"revenuePerShare"`
	NetIncomePerShare                      float64     `json:"netIncomePerShare"`
	OperatingCashFlowPerShare              float64     `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare                   float64     `json:"freeCashFlowPerShare"`
	CashPerShare                           float64     `json:"cashPerShare"`
	BookValuePerShare                      float64     `json:"bookValuePerShare"`
	TangibleBookValuePerShare              float64     `json:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare             float64     `json:"shareholdersEquityPerShare"`
	InterestDebtPerShare                   float64     `json:"interestDebtPerShare"`
	MarketCap                              float64     `json:"marketCap"`
	EnterpriseValue                        float64     `json:"enterpriseValue"`
	PeRatio                                float64     `json:"peRatio"`
	PriceToSalesRatio                      float64     `json:"priceToSalesRatio"`
	Pocfratio                              float64     `json:"pocfratio"`
	PfcfRatio                              float64     `json:"pfcfRatio"`
	PbRatio                                float64     `json:"pbRatio"`
	PtbRatio                               float64     `json:"ptbRatio"`
	EvToSales                              float64     `json:"evToSales"`
	EnterpriseValueOverEBITDA              float64     `json:"enterpriseValueOverEBITDA"`
	EvToOperatingCashFlow                  float64     `json:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64     `json:"evToFreeCashFlow"`
	EarningsYield                          float64     `json:"earningsYield"`
	FreeCashFlowYield                      float64     `json:"freeCashFlowYield"`
	DebtToEquity                           float64     `json:"debtToEquity"`
	DebtToAssets                           float64     `json:"debtToAssets"`
	NetDebtToEBITDA                        float64     `json:"netDebtToEBITDA"`
	CurrentRatio                           float64     `json:"currentRatio"`
	InterestCoverage                       float64     `json:"interestCoverage"`
	IncomeQuality                          float64     `json:"incomeQuality"`
	DividendYield                          float64     `json:"dividendYield"`
	PayoutRatio                            float64     `json:"payoutRatio"`
	SalesGeneralAndAdministrativeToRevenue int         `json:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDdevelopementToRevenue      float64     `json:"researchAndDdevelopementToRevenue"`
	IntangiblesToTotalAssets               int         `json:"intangiblesToTotalAssets"`
	CapexToOperatingCashFlow               float64     `json:"capexToOperatingCashFlow"`
	CapexToRevenue                         float64     `json:"capexToRevenue"`
	CapexToDepreciation                    float64     `json:"capexToDepreciation"`
	StockBasedCompensationToRevenue        float64     `json:"stockBasedCompensationToRevenue"`
	GrahamNumber                           float64     `json:"grahamNumber"`
	Roic                                   float64     `json:"roic"`
	ReturnOnTangibleAssets                 float64     `json:"returnOnTangibleAssets"`
	GrahamNetNet                           float64     `json:"grahamNetNet"`
	WorkingCapital                         int64       `json:"workingCapital"`
	TangibleAssetValue                     interface{} `json:"tangibleAssetValue"`
	NetCurrentAssetValue                   int64       `json:"netCurrentAssetValue"`
	InvestedCapital                        interface{} `json:"investedCapital"`
	AverageReceivables                     int64       `json:"averageReceivables"`
	AveragePayables                        int64       `json:"averagePayables"`
	AverageInventory                       int64       `json:"averageInventory"`
	DaysSalesOutstanding                   float64     `json:"daysSalesOutstanding"`
	DaysPayablesOutstanding                float64     `json:"daysPayablesOutstanding"`
	DaysOfInventoryOnHand                  float64     `json:"daysOfInventoryOnHand"`
	ReceivablesTurnover                    float64     `json:"receivablesTurnover"`
	PayablesTurnover                       float64     `json:"payablesTurnover"`
	InventoryTurnover                      float64     `json:"inventoryTurnover"`
	Roe                                    float64     `json:"roe"`
	CapexPerShare                          float64     `json:"capexPerShare"`
}

// KeyMetrics by Company
func KeyMetrics(apikey string) {
	responseData := getjsonresponse.GetResponse("https://financialmodelingprep.com/",
		"api/v3/key-metrics/",
		"AAPL",
		"demo")
	var km keyMetrics
	json.Unmarshal([]byte(responseData), &km)
	for _, value := range km {
		fmt.Println(value.Date)
	}

}
