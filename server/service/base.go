package service

import (
	"context"
	"github.com/go-terminal-server/common/constants"
	"gorm.io/gorm"
)

type baseService struct {
}

func (service baseService) Context(db *gorm.DB) context.Context {
	return context.WithValue(context.Background(), constants.DB, db)

}
