package redis


import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
)

var (
	RdsConn *redis.Client
	once	sync.Once
)


func init() {
	once.Do(func() {
		redisConfig, _ := beego.AppConfig.GetSection("redis")
		db_num, _ := strconv.Atoi(redisConfig["db_index"])
		RdsConn = redis.NewClient(&redis.Options{
			Addr:     redisConfig["address"],
			Password: redisConfig["password"],
			DB:       db_num,
		})
		_, err := RdsConn.Ping().Result()
		if err != nil {
			logs.Info("connect redis fail", err)
			panic(err)
		}
	})
}