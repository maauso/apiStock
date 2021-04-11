package financialmodelingprep

import (
	"apiStock/internal/domain/persistence"

	"github.com/zenthangplus/goccm"
)

//populator into the repository symbol and value.
func populator(dfc discountCashFlows, repo persistence.Repository) {
	c := goccm.New(10)
	for _, value := range dfc {
		c.Wait()
		go repo.PopulateData(value.Symbol, value.Dcf, c)
	}
	c.WaitAllDone()
}
