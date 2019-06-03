package user

import (
	"fmt"
	"github.com/kevenp/microbook/user-srv/basic/db"
	proto "github.com/kevenp/microbook/user-srv/proto/user"
	"github.com/micro/go-log"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {

	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	// 获取数据库
	o := db.GetDB()
	fmt.Printf("%p, %T\n", o, o)
	ret = &proto.User{}

	// 查询
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}