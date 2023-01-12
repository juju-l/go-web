package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go-web/api/vipex.cc/wx/v1"
	"go-web/internal/domain/dto"
	"go-web/util/ext"
	"net/http"
)

func reg()*runtime.ServeMux {
	var m *runtime.ServeMux = nil
	m = runtime.NewServeMux(
	runtime.WithErrorHandler(func(
	ctx context.Context,s *runtime.ServeMux,m runtime.Marshaler,w http.ResponseWriter,r *http.Request,err error,
	) {
	if err != nil {
	w.Header().Set("Content-Type","application/json");d,_:=m.Marshal(&dto.ResponseDto{ ErrCode:-2,Data:nil,ErrMsg:err.Error() });w.Write(d)
	}
	}),
	/*runtime.WithOutgoingHeaderMatcher(),*/
	runtime.WithMetadata(func(ctx context.Context,q *http.Request, /**/) metadata.MD {
	return metadata.Pairs("token",q.Header.Get("token"),/**/)
	}),
	)
	o := []grpc.DialOption{
	grpc.WithTransportCredentials(
	ext.Must(credentials.NewServerTLSFromFile("config/tls.crt","config/.tls.key",/**/)).(credentials.TransportCredentials),
	),
	}
	if e := v1.RegisterWxHandlerFromEndpoint(context.Background(),m,"i.vipex.cc:443",o,/**/);e != nil {
			panic(e)
	}
	return m
}

///
// other
//***

func init() {
	//
}