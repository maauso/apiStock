package apistockredis

import (
	"apiStock/internal/structure"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

//RedisRepository for Redis
type RedisRepository struct {
	pool redis.Conn
}

//NewRedisRepository for Redis
func NewRedisRepository(pool redis.Conn) *RedisRepository {
	return &RedisRepository{pool: pool}
}

//SaveData in Redis
func (r RedisRepository) SaveData(dfc structure.DiscountCashFlow) {
	_, err := r.pool.Do("SET", dfc.Symbol, dfc.Dcf)
	if err != nil {
		log.Fatal(err)
	}

	s, err := redis.String(r.pool.Do("GET", dfc.Symbol))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", s)
}
