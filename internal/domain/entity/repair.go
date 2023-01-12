package entity

//
///
type Repair struct {
	EntId string `gorm:"column:entid;type:varchar;comment:标识;primaryKey" json:"entid"`
	OpenId string `gorm:"<-:create;column:openid;type:varchar;comment:标识;primaryKey" json:"openid"`
	Repid string `gorm:"->;column:repid;type:varchar;comment:标识;primaryKey" json:"repid"`
	IsValid bool `gorm:"->;column:is_valid;type:varchar;comment:有效性" json:"is_valid"`
	Sid string `gorm:"column:sid;type:varchar;comment:简称" json:"sid"`
}

func init () {
	//
}