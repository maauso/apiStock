package main

import (
	"github.com/gomodule/redigo/redis"

	"apiStock/internal/application/selector"
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"apiStock/pkg/infrastructure/storage/apistockredis"
	"flag"
	"log"
	"os"
)

func main() {
	providerPtr := flag.String("provider", "", "Select the provider [ financialmodelingprep ] (Required)")
	metricPtr := flag.String("metric", "", "Metric { DiscountedCashFlow | HistoricalDiscountedCashFlow | KeyMetric | UnderValuatedCompanies };(Required)")
	companySymbolPtr := flag.String("company", "AAPL", "demo apikey")
	apikeyPtr := flag.String("apiKey", os.Getenv("apiKey"), "demo apikey")
	populateCompanies := flag.String("yes", "no", "Populate companies Symbol")
	flag.Parse()

	if *providerPtr == "" || *metricPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	repo := initializeRepo()
	params := structure.Arguments{Provider: *providerPtr, Metric: *metricPtr, Company: *companySymbolPtr, APIKey: *apikeyPtr, ListOfCompanies: *populateCompanies}
	selector.GetMetric(params, repo)
}

func initializeRepo() persistence.Repository {
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
