package selector

import (
	"apiStock/internal/application/financialmodelingprep"
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
)

// GetMetric Case between different providers
func GetMetric(p structure.Arguments, repo persistence.Repository) {
	if p.Provider == "financialmodelingprep" {
		financialmodelingprep.Fmg(structure.Arguments{Metric: p.Metric, Company: p.Company, APIKey: p.APIKey}, repo)
	}
}
