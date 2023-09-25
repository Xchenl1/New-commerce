package svc

import (
	"E-commerce_system/user/internal/config"
	"E-commerce_system/user/internal/middleware"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"strconv"
)

type ServiceContext struct {
	Config            config.Config
	DB                *gorm.DB
	JwtAuthMiddleware *middleware.JwtAuthMiddleware
	File              multipart.File
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		DB:                MysqlConn(c),
		JwtAuthMiddleware: middleware.NewJwtAuthMiddleware(),
	}
}

// MysqlConn 连接数据库
func MysqlConn(config config.Config) *gorm.DB {
	dsn := config.Mysql.User + ":" + config.Mysql.Password + "@tcp(" + config.Mysql.Host + ":" + strconv.Itoa(config.Mysql.Port) + ")/" + config.Mysql.Db + "?" + config.Mysql.Config
	fmt.Println(dsn)
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	//err = DB.AutoMigrate(
	//	model.Address{},
	//	model.Admin{},
	//	model.Carousel{},
	//	model.Cart{},
	//	model.Category{},
	//	model.Favorite{},
	//	model.Notice{},
	//	model.Order{},
	//	model.Product{},
	//	model.ProductImg{},
	//	model.User{},
	//)
	if err != nil {
		panic(err)
	}
	return DB
}
