package handler

import (
	"context"
	"fmt"
	"github.com/kevenp/microbook/user-srv/model/user"
	s "github.com/kevenp/microbook/user-srv/proto/user"
	"log"
)

type Service struct {

}

var (
	userService user.Service
)

func init()  {
	var err error
	userService, err = user.GetService()
	fmt.Printf("!!! %p, %T", userService, userService)
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *s.Request, rsp *s.Response) error {
	fmt.Printf("### %p, %T", userService, userService)
	userInfo, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &s.Error{
			Code:500,
			Detail:err.Error(),
		}
	}

	rsp.User = userInfo
	rsp.Success = true
	return err
}

