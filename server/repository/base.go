package repository

import (
	"context"
	"github.com/go-terminal-server/common"
	"github.com/go-terminal-server/common/constants"
	"gorm.io/gorm"
)

type baseRepository struct {
}

func (b *baseRepository) GetDB(c context.Context) *gorm.DB {
	db, ok := c.Value(constants.DB).(*gorm.DB)
	if !ok {
		return common.DB
	}
	return db
}
