package apistockredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/zenthangplus/goccm"

	"fmt"
	"log"
)

// RedisRepository for Redis
type RedisRepository struct {
	pool *redis.Pool
}

// NewRedisRepository for Redis
func NewRedisRepository(pool *redis.Pool) *RedisRepository {
	return &RedisRepository{pool: pool}
}

// PopulateData Save data in Redis
func (r *RedisRepository) PopulateData(Symbol string, DiscountCashFlow interface{}, c goccm.ConcurrencyManager) {
	conn := r.pool.Get()
	_, err := conn.Do("SET", Symbol, DiscountCashFlow)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("SAVING-> %v %v\n", Symbol, DiscountCashFlow)
	fmt.Printf("Worker %v: Started\n", Symbol)
	fmt.Printf("Worker %v: Finished\n", Symbol)
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Done()
}

// PopulateUnderValuatedCompanies to be implemented soon
func (r *RedisRepository) PopulateUnderValuatedCompanies(Symbol string, Percentage float64, c goccm.ConcurrencyManager) {
	conn := r.pool.Get()
	_, err := conn.Do("SET", "UnderValuated"+Symbol, Percentage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Worker %v: Started\n", Symbol)
	fmt.Printf("Worker %v: Finished\n", Symbol)
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Done()
}

// GetTotalCompanies get company value from repository
func (r *RedisRepository) GetTotalCompanies(listOfElements string, companiesChannel chan<- string, c goccm.ConcurrencyManager) {
	conn := r.pool.Get()
	_, err := redis.Float64(conn.Do("GET", listOfElements))
	if err == nil {
	} else {
		companiesChannel <- listOfElements
	}
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Elementos en la cola %v \n", len(companiesChannel))
	defer c.Done()

}
