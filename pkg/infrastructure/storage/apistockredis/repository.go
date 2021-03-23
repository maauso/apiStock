package apistockredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/zenthangplus/goccm"

	"fmt"
	"log"
)

//RedisRepository for Redis
type RedisRepository struct {
	pool redis.Pool
}

func NewRedisRepository(pool redis.Pool) *RedisRepository {
	return &RedisRepository{pool: pool}
}

//NewRedisRepository for Redis

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

func (r RedisRepository) GetData() []string {
	conn := r.pool.Get()
	keys, err := redis.Strings(conn.Do("KEYS", "*"))
	if err != nil {
		log.Fatal(err)
	}
	return keys
}
