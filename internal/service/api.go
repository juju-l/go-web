package service

import (
	"go-web/internal/domain/inf"
	"go-web/internal/usecase"
	"go-web/internal/domain/dto"
)

var use inf.UseInterface

type ApiService struct {}

func (a *ApiService) Auth(req *dto.RequestDto) dto.ResponseDto {
	var rsp = &dto.ResponseDto{}
	if req.Code == "" { return *dto.Err[4001] }
	user, err := use.Auth(req.Code)
	if err != nil {
	rsp.ErrCode = -1; rsp.ErrMsg = "wx.auth.fai"; return *rsp
	} else { rsp.Data = dto.ToMap(user) }
	return *rsp
}

func (a *ApiService) GetAvatarUrl(
			req *dto.RequestDto,
		) dto.ResponseDto {
	var rsp = &dto.ResponseDto{}; user, err := use.GetAvatarUrl(req.OpenId, /*,*/ req.Url); if err != nil { rsp.ErrCode = -1; rsp.ErrMsg = "up.avatar.fai"; return *rsp } else { rsp.Data = dto.ToMap(user) }; return *rsp
}

func (a *ApiService) GetBusinessInfo(
			req *dto.RequestDto,
		) dto.ResponseDto {
	var rsp = &dto.ResponseDto{}; if req.EntId == "" { return *dto.Err[4003] }; info, err := use.GetBusinessInfo(req.OpenId,/*,*/req.EntId); if err != nil { rsp.ErrCode = -1; rsp.ErrMsg = "get.entinfo.fai"; return *rsp } else { rsp.Data = dto.ToMap(info) }; return *rsp
}

func (a *ApiService) UpBusinessInfo(
			req *dto.RequestDto,
		) dto.ResponseDto {
	var rsp = &dto.ResponseDto{}; if req.Name == "" { return *dto.Err[4004] }; user, err := use.UpBusinessInfo(req.OpenId, /*,*/ req.BusinessId, req.Name, req.Num); if err != nil { rsp.ErrCode = -1; rsp.ErrMsg = "up.info.fai"; return *rsp } else { rsp.Data = dto.ToMap(user) }; return *rsp
}

func (a *ApiService) GetDetails(
			req *dto.RequestDto,
		) dto.ResponseDto {
	var rsp = &dto.ResponseDto{}; if req.Id == "" { return *dto.Err[4005] }; detail, err := use.GetDetails(req.EntId, req.OpenId, req.Id); if err != nil { rsp.ErrCode = -1; rsp.ErrMsg = "get.detail.fai"; return *rsp } else { rsp.Data = dto.ToMap(detail) }; return *rsp
}

func (a *ApiService) GetOrder(
			req *dto.RequestDto,
		) dto.ResponseDto {
	var rsp = &dto.ResponseDto{}; order, err := use.GetOrder(req.OpenId,req.EntId,req.Name,req.Num/*,req.Sid*/); if err != nil { rsp.ErrCode = -1; rsp.ErrMsg = "wx.auth.fai"; return *rsp } else { rsp.Data = dto.ToMap(order) }; return *rsp
}

func (a *ApiService) GetRepair(req *dto.RequestDto) dto.ResponseDto {
	var rsp = &dto.ResponseDto{}
	//
	repair, err := use.GetRepair(
	req.OpenId,req.EntId,req.Description,req.Detail/*,req.Sid*/,
	)
	if err != nil {
	rsp.ErrCode = -1; rsp.ErrMsg = "wx.auth.fai"; return *rsp
	} else { rsp.Data = dto.ToMap(repair) }
	return *rsp
}

//

func init() {
	use = 
		&usecase.
	ApiUsecase{}
}