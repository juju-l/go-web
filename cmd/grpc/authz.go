package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"go-web/api/vipex.cc/wx/v1"
	"go-web/util/ext"
	//
)

//
///
func Authz(c context.Context, q interface{}, i *grpc.UnaryServerInfo, h grpc.UnaryHandler, /**/)(interface{},error) {
	if i.FullMethod != "/v1.Wx/Auth" {
			t := metadata.ValueFromIncomingContext(c,"token")
			if len(t) > 0 { if s,_ := ext.ValidateTkn(t[0]); s { return h(c,q,/**/) } }
			return &v1.WxResponse{ ErrCode:-1,Data:nil,ErrMsg:"401" },nil
	} else {
			r,e := h(c,q,/**/)
			if e == nil {
					if d,o := r.(*v1.WxResponse); o {
							if d.Data != nil {
			u := d.Data.AsMap()
							w,_ := ext.New(u["entid"].(string), u["is_authorization"].(bool), u["openid"].(string))
			s,_ := ext.GenToken(w);j := struct { Type string `json:"type"`; Token string `json:"token"`; Exp int `json:"exp"` }{"jwt",s,7200};b,_ := json.Marshal(j)
							d.Data.UnmarshalJSON(b)
			return d, nil
							}
					}
			}
			return r,e
	}
}

//

func init() {
	//
}