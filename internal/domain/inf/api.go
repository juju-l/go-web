package inf

import (
	"go-web/internal/domain/dto"
)

//

type ApiInterface interface {
	Auth(req *dto.RequestDto) dto.ResponseDto
	GetAvatarUrl(req *dto.RequestDto) dto.ResponseDto
	GetBusinessInfo(req *dto.RequestDto) dto.ResponseDto
	UpBusinessInfo(req *dto.RequestDto) dto.ResponseDto
	GetDetails(req *dto.RequestDto) dto.ResponseDto
	GetOrder(req *dto.RequestDto) dto.ResponseDto
	GetRepair(req *dto.RequestDto) dto.ResponseDto
}

func init() {
	//
}