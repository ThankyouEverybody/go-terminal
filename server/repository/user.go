package repository

import (
	"context"
	"github.com/go-terminal-server/model/orm"
)

var UserRepository = new(userRepository)

type userRepository struct {
	baseRepository
}

func (r userRepository) FindAll(c context.Context) (u []orm.User, err error) {
	err = r.GetDB(c).Find(&u).Error
	return
}

func (r userRepository) Create(c context.Context, u *orm.User) error {
	return r.GetDB(c).Create(u).Error
}
