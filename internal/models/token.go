package models

import (
	"time"

	"gorm.io/datatypes"
)

type CloudToken struct {
	ID          int64             `gorm:"primaryKey" json:"id"`
	Name        string            `gorm:"column:name;type:varchar(255);not null" json:"name"`
	AccessToken string            `gorm:"column:access_token;type:varchar(255);not null" json:"accessToken"`
	ExpiresIn   int64             `gorm:"column:expires_in;type:bigint(20);not null" json:"expiresIn"`
	Status      int8              `gorm:"column:status;type:tinyint(1);default:1" json:"status"`        // 状态 1:正常 2: 登录失败
	LoginType   int8              `gorm:"column:login_type;type:tinyint(1);default:1" json:"loginType"` // 1: 扫码登录 2: 密码登录
	Username    string            `gorm:"column:username;type:varchar(255);not null;default:''" json:"username"`
	Password    string            `gorm:"column:password;type:varchar(255);not null;default:''" json:"-"`
	Addition    datatypes.JSONMap `gorm:"column:addition;type:json;default:'{}'" json:"addition"` // 附属参数
	CreatedAt   time.Time         `gorm:"column:created_at;autoCreateTime;type:datetime;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time         `gorm:"column:updated_at;autoUpdateTime;type:datetime;default:CURRENT_TIMESTAMP;on update:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (c *CloudToken) TableName() string {
	return "cloud_tokens"
}

const (
	CloudTokenAdditionAutoLoginResultKey = "auto_login_result"
	CloudTokenAdditionAutoLoginTimes     = "auto_login_times"
)

const (
	LoginTypeScan = iota + 1
	LoginTypePassword
)
