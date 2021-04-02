package structure

// Arguments structure to initial arguments
type Arguments struct {
	Provider             *string
	Metric               *string
	Company              *string
	APIKey               *string
	ListOfCompanies      *string
	PercentageOfGrowth   *float64
	UnderValuedArguments UnderValuedArguments
}

// NewArguments constructor
func NewArguments(provider *string, metric *string, company *string, APIKey *string, listOfCompanies *string, percentageOfGrowth *float64, underValuedArguments UnderValuedArguments) *Arguments {
	return &Arguments{Provider: provider, Metric: metric, Company: company, APIKey: APIKey, ListOfCompanies: listOfCompanies, PercentageOfGrowth: percentageOfGrowth, UnderValuedArguments: underValuedArguments}
}

// UnderValuedArguments struct
type UnderValuedArguments struct {
	MarketCap *string
	Beta      *string
	Sector    *string
}

// NewUnderValuedArguments constructor
func NewUnderValuedArguments(marketCap *string, beta *string, sector *string) *UnderValuedArguments {
	return &UnderValuedArguments{MarketCap: marketCap, Beta: beta, Sector: sector}
}
