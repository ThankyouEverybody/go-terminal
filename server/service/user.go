package service

import (
	"context"
	"github.com/go-terminal-server/common/constants"
	"github.com/go-terminal-server/model"
	"github.com/go-terminal-server/model/orm"
	"github.com/go-terminal-server/repository"
)

var UserService = new(userService)

const SuperAdminID = `sdsdsdsdsdsdsdssdsdsdsdsdsdsd`

type userService struct {
	baseService
}

func (s userService) InitUser() error {
	users, err := repository.UserRepository.FindAll(context.TODO())
	if err != nil {
		return err
	}
	if len(users) <= 0 {

		//创建默认用户
		user := orm.User{
			ID:        SuperAdminID,
			UserName:  "超级管理员",
			Password:  "admin",
			LoginName: "admin",
			Mail:      "admin@goTerminal.com",
			Phone:     "13888888888",
			Status:    constants.N0,
			Type:      constants.SuperAdminUser,
		}
		user.CreateUid = SuperAdminID
		user.CreateTime = model.NowJsonTime()
		return repository.UserRepository.Create(context.TODO(), &user)
	}

	return nil
}
