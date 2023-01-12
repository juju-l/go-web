package inf

import (
	"go-web/internal/domain/entity"
)

//

type UseInterface interface {
	Auth(
			code string,
		) (*entity.User, error)
	GetAvatarUrl(openid,/*,*/url string) (*entity.User, error)
	GetBusinessInfo(openid,/*,*/entid string) (interface{}, error)
	UpBusinessInfo(openid,/*,*/businessid,name,num string) (*entity.Ent, error)
	GetDetails(entId,openid,id string) (map[string]interface{}, error)
	GetOrder(openid,entid,name,num/*,sid*/ string) (*[]entity.Detail, error)
	GetRepair(
			openid,entid,description,details/*,sid*/ string,
		) (*[]entity.Detail, error)
}

func init() {
	//
}