package structure

// Arguments structure to initial arguments
type Arguments struct {
	Provider string
	Metric   string
	Company  string
	APIKey   string
}

//DiscountCashFlow Json structure
type DiscountCashFlow struct {
	Symbol     string  `json:"symbol"`
	Date       string  `json:"date"`
	Dcf        float64 `json:"dcf"`
	StockPrice float64 `json:"Stock Price"`
}
