package main

import (
	"apiStock/internal/application/selector"
	"apiStock/internal/domain/derefstring"
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"apiStock/pkg/infrastructure/storage/apistockredis"
	"flag"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

func main() {
	providerPtr := derefstring.DerefString(flag.String("provider", "financialmodelingprep", "Select the provider [ financialmodelingprep ] (Required)"))
	metricPtr := derefstring.DerefString(flag.String("metric", "DiscountedCashFlow", "Metric { DiscountedCashFlow | HistoricalDiscountedCashFlow | KeyMetric };(Required)"))
	companySymbolPtr := derefstring.DerefString(flag.String("company", "AAPL", "demo apikey"))
	apikeyPtr := derefstring.DerefString(flag.String("apikey", "demo", "demo apikey"))

	flag.Parse()

	if providerPtr == "" || metricPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	repo := initializeRepo()
	params := structure.Arguments{Provider: providerPtr, Metric: metricPtr, Company: companySymbolPtr, APIKey: apikeyPtr}
	selector.GetMetric(params, repo)
}

//initializeRepo
func initializeRepo() persistence.Repository {
	repo := newRedisRepository()
	return repo
}

func newRedisRepository() persistence.Repository {
	c, err := redis.Dial("tcp",
		"redis-16211.c238.us-central1-2.gce.cloud.redislabs.com:16211",
		redis.DialPassword(""),
	)

	if err != nil {
		log.Fatal(err)
	}

	return apistockredis.NewRedisRepository(c)
}
