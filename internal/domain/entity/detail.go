package entity

import "time"

//

type Detail struct {
	EntId string `gorm:"column:entid;type:varchar;comment:标识;primaryKey" json:"entid"`
	OpenId string `gorm:"<-:create;column:openid;type:varchar;comment:标识;primaryKey" json:"openid"`
	Id string `gorm:"column:id;type:varchar;comment:标识;primaryKey" json:"id"`
	Detail string `gorm:"column:detail;type:varchar;comment:" json:"detail"`
	TimeSp time.Time `gorm:"->;column:timesp;type:varchar;comment:" json:"timesp"`
	IsValid bool `gorm:"->;column:is_valid;type:varchar;comment:有效性" json:"is_valid"`
	Sid string `gorm:"column:sid;type:varchar;comment:简称" json:"sid"`
}

func init () {
	//
}