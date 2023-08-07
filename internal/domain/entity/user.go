package entity

//
///
type User struct {
	OpenId string `gorm:"column:openid;type:varchar;comment:标识;primaryKey" json:"openid"`
	SessionKey string `gorm:"column:session_key;type:varchar;comment:会话密钥" json:"session_key"`
	UnionId string `gorm:"column:unionid;type:varchar;comment:标识" json:"unionid"`
	AvatarUrl string `gorm:"column:avatar_url;type:varchar;comment:头像" json:"avatarUrl"`
	IsAuthorization bool `gorm:"column:is_authorization;type:varchar;comment:授权" json:"is_authorization"`
	EntId string `gorm:"->;column:entid;type:varchar;comment:企业标识" json:"entid"`
	IsValid bool `gorm:"->;column:is_valid;type:varchar;comment:有效性" json:"is_valid"`
}

func init () {
	//
}