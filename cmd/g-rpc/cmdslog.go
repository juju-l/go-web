package main

import (
	"context"
	"google.golang.org/grpc"
	"go-web/api/vipex.cc/wx/v1"
	//"encoding/json"
	"fmt"
)

//
///
func Log(
					c context.Context, q interface{}, i *grpc.UnaryServerInfo, h grpc.UnaryHandler, /**/
)(interface{},error) {
	r := q.(*v1.WxRequest)
	fmt.Print(`grpcurl -insecure \
-d '{
	"code": "`+r.Code+`",
	"openid": "xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	"entid": "001d6fdb-79b5-41ff-af29-99a0e8618c37",
	"url": "",
	"businessid": "",
	"name": "",
	"num": "",
	"type": "",
	"description": "",
	"detail": "",
	"id": "8d3df1e1-9891-4d23-8435-a62c3e3bf6a1"
}' \
-proto wxService.proto \
-v -H token:\ `+"`"+`echo -n|grpcurl -insecure -d '{"code":"debug"}' -proto wxService.proto 127.0.0.1:443 v1.Wx/Auth|jq .data.token -r`+"`"+` -emit-defaults \
127.0.0.1:443 \
v1.Wx/\
$method`)
	return h(c,q,/**/)
}

//

func init() {
	//
}