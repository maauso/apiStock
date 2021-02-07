package selector

import (
	"apiStock/internal/application/financialmodelingprep"
	"apiStock/pkg/infrastructure/http/arguments"
)

// GetMetric Case between different providers
func GetMetric(p arguments.Arguments) {
	if p.Provider == "financialmodelingprep" {
		financialmodelingprep.Fmg(arguments.Arguments{Metric: p.Metric, Company: p.Company, APIKey: p.APIKey})
	}
}
