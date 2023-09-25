package svc

import (
	"E-commerce_system/order/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	//Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     MysqlConn(c),
		//	Redis:  RedisConn(c),
	}
}

// MysqlConn 连接数据库
func MysqlConn(config config.Config) *gorm.DB {
	dsn := config.Mysql.User + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Host + ":" + strconv.Itoa(config.Mysql.Port) + ")/" + config.Mysql.Db + "?" + config.Mysql.Config
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("MySQL连接成功！")
	//DB.AutoMigrate(model.Product{}, model.Favorite{})
	return DB
}

func RedisConn(config config.Config) *redis.Client {
	RedisConn0 := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port),
		Password: config.Redis.Password, // Redis 未设置密码时为空
		DB:       0,                     // 使用默认数据库0
	})
	_, err := RedisConn0.Ping(RedisConn0.Context()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis连接成功！")
	return RedisConn0
}
