package apistockredis

import (
	"apiStock/internal/structure"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisRepository struct {
	pool redis.Conn
}

func NewRedisRepository(pool redis.Conn) *RedisRepository {
	return &RedisRepository{pool: pool}
}

func (r RedisRepository) SaveData(dfc structure.DiscountCashFlow) {
	r.pool.Do("SET", dfc.Symbol, dfc.Dcf)
	s, _ := redis.String(r.pool.Do("GET", dfc.Symbol))
	fmt.Printf("%#v\n", s)
}
