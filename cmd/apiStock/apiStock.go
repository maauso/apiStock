package main

import (
	"apiStock/internal/application/selector"
	"apiStock/internal/domain/derefstring"
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"apiStock/pkg/infrastructure/storage/apistockredis"
	"flag"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

func main() {
	providerPtr := derefstring.DerefString(flag.String("provider", "financialmodelingprep", "Select the provider [ financialmodelingprep ] (Required)"))
	metricPtr := derefstring.DerefString(flag.String("metric", "UnderValuatedCompanies", "Metric { DiscountedCashFlow | HistoricalDiscountedCashFlow | KeyMetric | UnderValuatedCompanies };(Required)"))
	companySymbolPtr := derefstring.DerefString(flag.String("company", "AAPL", "demo apikey"))
	apikeyPtr := derefstring.DerefString(flag.String("apiKey", os.Getenv("apiKey"), "demo apikey"))
	populateCompanies := derefstring.DerefString(flag.String("yes", "yes", "Populate companies Symbol"))
	flag.Parse()

	if providerPtr == "" || metricPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	repo := initializeRepoPool()
	params := structure.Arguments{Provider: providerPtr, Metric: metricPtr, Company: companySymbolPtr, APIKey: apikeyPtr, ListOfCompanies: populateCompanies}
	selector.GetMetric(params, repo)
}

func initializeRepoPool() persistence.Repository {
	repo := newRedisRepository()
	return repo
}

func newRedisRepository() persistence.Repository {
	pool := redis.Pool{
		MaxActive:   1,
		MaxIdle:     0,
		IdleTimeout: 1,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp",
				"redis-16211.c238.us-central1-2.gce.cloud.redislabs.com:16211",
				redis.DialPassword(os.Getenv("RedisLabKey")),
			)
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
	return apistockredis.NewRedisRepository(pool)
}
