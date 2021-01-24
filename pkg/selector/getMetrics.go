package selector

import (
	structs "apiStock/model"
	"apiStock/pkg/financialmodelingprep"
)

// GetMetric Case between different providers
func GetMetric(p structs.Arguments) {
	if p.Provider == "financialmodelingprep" {
		financialmodelingprep.Fmg(structs.Arguments{Metric: p.Metric, Company: p.Company, APIKey: p.APIKey})
	}
}
