package main

import (
	"github.com/gomodule/redigo/redis"

	"apiStock/internal/application/selector"
	"apiStock/internal/domain/persistence"
	"apiStock/internal/structure"
	"apiStock/pkg/infrastructure/storage/apistockredis"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	providerPtr := flag.String("provider", "financialmodelingprep", "Select the provider [ financialmodelingprep ] (Required)")
	metricPtr := flag.String("metric", "", "Metric { DiscountedCashFlow | UnderValuatedCompanies };(Required)")
	companySymbolPtr := flag.String("company", "MSFT", "MSFT")
	apikeyPtr := flag.String("apiKey", os.Getenv("apiKey"), "financialmodelingprep apikey (Required)")
	marketCapPtr := flag.String("marketCap", "1000000000", "Market Capitalization ")
	betaPtr := flag.String("beta", "1", "Company Beta")
	sectorPtr := flag.String("sector", "Technology", "Company Sector: Consumer Cyclical - Energy - Technology - Industrials - Financial Services - Basic Materials - Communication Services - Consumer Defensive - Healthcare - Real Estate - Utilities - Industrial Goods - Financial - Services - Conglomerates ")
	percentageOfGrowthPtr := flag.Float64("percentageOfGrowth", 10, "Set the minimum growth tha the company should have")
	flag.Parse()

	if *apikeyPtr == "" || *metricPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *metricPtr == "UnderValuatedCompanies" {
		fmt.Println("Info: Be sure that you set marketCap, beta, sector")
	}

	repo := initializeRepo()
	underValuesParams := structure.NewUnderValuedArguments(marketCapPtr, betaPtr, sectorPtr)
	params := structure.NewArguments(providerPtr, metricPtr, companySymbolPtr, apikeyPtr, percentageOfGrowthPtr, *underValuesParams)
	selector.GetMetric(params, repo)
}

func initializeRepo() persistence.Repository {
	repo := newRedisRepository()
	return repo
}

func newRedisRepository() persistence.Repository {
	pool := redis.Pool{
		MaxActive:   10,
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
