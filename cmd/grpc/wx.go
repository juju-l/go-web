package main

import (
	"context"
	//
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/grpc/reflection"
	"go-web/api/vipex.cc/wx/v1"
	"go-web/internal/domain/inf"
	"go-web/internal/service"
	"go-web/internal/domain/dto"
	"google.golang.org/grpc/credentials"
	"go-web/util/ext"
	//#"openapi/v3"
	"net"
	//#"net/http"
)

var svc inf.ApiInterface

type wxService struct{}

			func (rpc *wxService) Auth(ctx context.Context, req *v1.WxRequest) (*v1.WxResponse, error) {return mapRp( svc.Auth( mapRq(req) ) ), nil} /**/ //
			func (rpc *wxService) GetAvatarUrl(ctx context.Context, req *v1.WxRequest) (*v1.WxResponse, error) {return mapRp( svc.GetAvatarUrl( mapRq(req) ) ), nil} /**/ //
			func (rpc *wxService) GetBusinessInfo(ctx context.Context, req *v1.WxRequest) (*v1.WxResponse, error) {return mapRp( svc.GetBusinessInfo( mapRq(req) ) ), nil} /**/ //
			func (rpc *wxService) UpBusinessInfo(ctx context.Context, req *v1.WxRequest) (*v1.WxResponse, error) {return mapRp( svc.UpBusinessInfo( mapRq(req) ) ), nil} /**/ //
			func (rpc *wxService) GetDetails(ctx context.Context, req *v1.WxRequest) (*v1.WxResponse, error) {return mapRp( svc.GetDetails( mapRq(req) ) ), nil} /**/ //
			func (rpc *wxService) GetOrder(ctx context.Context, req *v1.WxRequest) (*v1.WxResponse, error) {return mapRp( svc.GetOrder( mapRq(req) ) ), nil} /**/ //
			func (rpc *wxService) GetRepair(ctx context.Context, req *v1.WxRequest) (*v1.WxResponse, error) {return mapRp( svc.GetRepair( mapRq(req) ) ), nil} /**/ //

func main() {
	l, _ := net.Listen(
					"tcp", ":443", /**/
	)
	s := grpc.NewServer(
	grpc.UnaryInterceptor(
			Authz,
	),
		grpc.Creds(ext.Must(
		credentials.
		NewServerTLSFromFile(
			"config/tls.crt",
			"config/.tls.key",
			///
			)).
		(credentials.
	TransportCredentials)),
	//grpc.StreamInterceptor(),
	)
	v1.RegisterWxServer(
			s, &wxService{}, /*,*/
	)
	//
		reflection.Register(s)
			//
	if e:=s.Serve(l);e!=nil {
			panic( e )
	}
			//
	/*#if e:=http.ServeTLS(
			l,
	grpcHandlerFunc(
	s, v3.OpenapiV3(), reg(),
	),
			"config/tls.crt",
			"config/.tls.key",
			///
	);e!=nil {
					panic( e )/////
	}#*/
}

func mapRq(
			req *v1.WxRequest,
	) *dto.RequestDto {
 return &dto.RequestDto {
			Code: req.Code, OpenId: req.Openid, EntId: req.Entid, Url: req.Url, BusinessId: req.Businessid, Name: req.Name, Num: req.Num, Type: req.Type, Description: req.Description, Detail: req.Detail, /*OrdId: req.OrdId, Repid: req.Repid,*/ Id: req.Id,
 }
}

func mapRp(
			rsp dto.ResponseDto,
		) *v1.WxResponse {
	r := &v1.WxResponse{}
	if rsp.ErrCode == 0 { d := &structpb.Struct{};b,e := json.Marshal(rsp.Data);if e != nil {return r};d.UnmarshalJSON( b );r.Data = d } else {
			r.ErrCode=int32(rsp.ErrCode)
			/*;*/
			r.ErrMsg=rsp.ErrMsg
	}
	return r
}

func init() {
	//
	svc = 
			&service.
	ApiService{}
	//
}