package subscriber

import (
	"context"
	"github.com/micro/go-log"

	user "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/proto/user"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *user.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
