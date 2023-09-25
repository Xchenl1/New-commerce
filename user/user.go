package main

import (
	"E-commerce_system/user/internal/config"
	"E-commerce_system/user/internal/handler"
	"E-commerce_system/user/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config

	conf.MustLoad(*configFile, &c)           //加载yaml文件到c中
	server := rest.MustNewServer(c.RestConf) //通过config下的restconf来注册服务
	defer server.Stop()

	//处理预检请求
	server.Use(corsMiddleware)

	//静态访问图片
	ctx := svc.NewServiceContext(c)       //添加到ServiceContext
	handler.RegisterHandlers(server, ctx) //通过server和ctx注册路由

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST,GET,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}
