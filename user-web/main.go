package main

import (
	"fmt"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-web"
	"github.com/kevenp/microbook/user-srv/basic/config"
	"github.com/kevenp/microbook/user-web/handler"
	"time"
)

func main()  {
	micReg := consul.NewRegistry(registryOptions)

	service := web.NewService(
		// 后面两个web，第一个是指是web类型的服务，第二个是服务自身的名字
		web.Name("mu.micro.book.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	// 注册登录接口
	service.HandleFunc("/user/login", handler.Login)
	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}




