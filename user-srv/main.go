package main

import (
	"fmt"
	"github.com/kevenp/microbook/user-srv/basic/config"
	"github.com/kevenp/microbook/user-srv/handler"
	"github.com/kevenp/microbook/user-srv/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"time"
)

func main() {

	//使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)
	srv := new(handler.Service)
	service.Init()
	err := mu_micro_book_srv_user.RegisterUserHandler(service.Server(), srv)

	if err != nil {
		log.Fatal(err)
	}
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options)  {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}