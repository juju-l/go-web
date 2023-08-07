package entity

//
///
type Ent struct {
	EntId string `gorm:"column:entid;type:varchar;comment:标识;primaryKey" json:"entid"`
	OpenId string `gorm:"<-:create;column:openid;type:varchar;comment:标识;primaryKey" json:"openid"`
	Name string `gorm:"column:name;type:varchar;comment:企业全称" json:"name"`
	Num string `gorm:"column:num;type:varchar;comment:企业税号" json:"num"`
	Oth string `gorm:"column:oth;type:varchar;comment:---备用字段---" json:"oth"`
	IsValid bool `gorm:"->;column:is_valid;type:varchar;comment:有效性" json:"is_valid"`
	Sid string `gorm:"column:sid;type:varchar;comment:简称/*;primaryKey*/" json:"sid"`
}

func init () {
	//
}