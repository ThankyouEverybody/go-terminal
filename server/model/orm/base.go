package orm

import (
	"github.com/go-terminal-server/model"
)

type base struct {
	CreateTime  model.JsonTime `gorm:"type:datetime;comment:创建时间;<-:create" json:"createTime"`
	CreateUid   string         `gorm:"type:varchar(36);not null;comment:创建人id;<-:create" json:"createUid"`
	UpdatedTime model.JsonTime `gorm:"type:datetime;comment:修改时间;<-:update" json:"updateTime"`
	UpdateUid   string         `gorm:"type:varchar(36);not null;comment:更新人id;<-:update" json:"updateUid"`
}
