package apistockredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/zenthangplus/goccm"

	"fmt"
	"log"
)

// RedisRepository for Redis
type RedisRepository struct {
	pool redis.Pool
}

// NewRedisRepository for Redis
func NewRedisRepository(pool redis.Pool) *RedisRepository {
	return &RedisRepository{pool: pool}
}

// PopulateData Save data in Redis
func (r RedisRepository) PopulateData(Sector, Symbol, CompanyName string, c goccm.ConcurrencyManager) {
	conn := r.pool.Get()
	exists, _ := redis.Bool(conn.Do("EXISTS", Sector+" "+Symbol))

	if !exists {
		_, err := conn.Do("SET", Sector+" "+Symbol, CompanyName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("SAVING->" + Sector + " " + Symbol + ":" + CompanyName)
	}
	fmt.Printf("Worker %v: Started\n", Symbol)
	fmt.Printf("Worker %v: Finished\n", Symbol)
	err := conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Done()
}

// PopulateUnderValuatedCompanies to be implemented soon
func (r RedisRepository) PopulateUnderValuatedCompanies(Symbol string, Percentage float64, c goccm.ConcurrencyManager) {
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

// GetData from Redis
func (r RedisRepository) GetData() []string {
	conn := r.pool.Get()
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		log.Fatal("Error in GetData function: ", err)
	}
	return keys
}
