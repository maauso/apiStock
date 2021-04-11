package selector

import (
	"apiStock/internal/application/financialmodelingprep"
	"apiStock/internal/arguments"
	"apiStock/internal/domain/persistence"
)

// GetMetric Case between different providers
func GetMetric(arguments *arguments.Arguments, repo persistence.Repository) {
	if *arguments.Provider == "financialmodelingprep" {
		financialmodelingprep.Fmg(arguments, repo)
	}
}

// La idea es que a la funcion final le lleguen unicamente los parametros que le hacen falta, aunque le llegue el string directamente.
