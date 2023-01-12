package ext

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"go-web/api/vipex.cc/wx/v1"
	"go-web/internal/domain/entity"
	"go-web/internal/domain/dto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"github.com/gin-gonic/gin"
	//"strconv"
	"strings"
)

// type tknHeader struct {
// 	Token string `header:"token"`
// }

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.URL.Path
		var j=struct { ///
			Type string `json:"type"`
			Token string `json:"token"`
			Exp int `json:"exp"`
		}{
		"jwt","",7200,
		}
		// token := &tknHeader{}
		//if uri != "/api/auth" {
		if !strings.Contains(strings.ToLower(uri),"/auth") {
		//if uri!="/api/auth"||uri!="/v1.Wx/Auth" {
			//|if c.GetHeader("token") != "" {
					if s,_ :=
			ValidateTkn(c.GetHeader("token"));
				s == true||strings.HasPrefix(
			//reflection.Register( grpcSvc )
			///grpc.reflection.v1alpha.ServerReflection/ServerReflectionInfo
					uri,/*Reflect->,*/"/grpc",
					)||
					strings.
					Contains(
					uri,/*swagger->,*/"/openapi",
					) {
			c.Next();/*must->;*/return
					} else {
					g := "application/grpc"
					if c.ContentType()==g {
			c.Status(401)
					//|c.Writer.Header().Set("Content-Type",g,/**/)
					//c.Writer.Header().Add("Trailer","Grpc-Status",/**/)
					//c.Writer.Header().Add("Trailer","Grpc-Status-Details-Bin",/**/)
					//c.Writer.Header().Add("Trailer","Grpc-Message",/**/)
					//c.Writer.Header().Set("Grpc-Status","401",/**/)
					o := c.Writer
			r := &rspBody{
					rsb: bytes.NewBufferString(""),
					ResponseWriter: c.Writer,
					// str &""
			}
					c.Writer = r
					c.Next() // 否则rsp未初始化
					c.Writer = o
					p:=&v1.WxResponse{
			ErrCode:-1,Data:nil,ErrMsg:"401",
					}
					n,_:=proto.Marshal(p)
					q:=make([]byte, 5)
					binary.BigEndian.PutUint32(
							q[1:], uint32(len(n)),
					)
					for _,v:=range n{
							q=append(q,v)
					}
					n=q
					c.Writer.Write(n)
					c.Abort()
					} else {
					c.AbortWithStatusJSON(
			401,&dto.ResponseDto{-1,nil,"401"},//
					)
					}
					return
					}
			//|}
		} else {
					o := c.Writer
			r := &rspBody{
					rsb: bytes.NewBufferString(""),
					ResponseWriter: c.Writer,
					// str &""
			}
					c.Writer = r
					c.Next()
					d := &dto.ResponseDto{}
					p := &v1.WxResponse{
							Data:&structpb.Struct{},
					}
					u := &entity.User{}
			if c.Writer.Status() == 200 {
					b := r.rsb.Bytes()
					///
					if e := proto.Unmarshal(
					r.rsb.Bytes()[5:], p, /**/
					);e == nil {
					b,_ = json.Marshal(p)
					}
							if json.Unmarshal(b,d);d.Data != nil {
									if t,e := json.Marshal(d.Data);e == nil {
											if e := json.Unmarshal(t, u);e == nil {
													if w,e := New(u.EntId, u.IsAuthorization, u.OpenId);e == nil {
															if s,e := GenToken(w);e == nil {j.Token=s;d.Data=dto.ToMap(j);//
																	if n,e := json.Marshal(d);e == nil {
																			//
																			c.Writer = o;/*c.Header("token",s);c.Header("exp",strconv.FormatInt(w.ExpiresAt.Unix(),10));*//*p=mapRp(d);*/if uri=="/v1.Wx/Auth" {z,_:=json.Marshal(j);p.Data.UnmarshalJSON(z);p.ErrCode=0;p.ErrMsg="";n,_=proto.Marshal(p);q:=make([]byte, 5);binary.BigEndian.PutUint32(q[1:], uint32(len(n)));for _,v:=range n{q=append(q,v)};n=q};c.Writer.Write(n);return
																			//
																	}
															}
													}
											}
									}
							}
			}
			c.Writer = o
			c.Writer.Write(
					r.rsb.Bytes(),
			)
			// default 
		}
		///
		//
		//|c.AbortWithStatus/*JSON*/(
	//|401,//&dto.ResponseDto{-1,nil,"401"},//
		//|)
		// c.Next()
	}
}

type rspBody struct {
	rsb *bytes.Buffer
	gin.ResponseWriter
	// str *string
}

func (r *rspBody) Write(
			b []byte,
		) (int, error) {
	return r.rsb.Write(b)
	///i, e := r.ResponseWriter.Write(b)
	///return i, e
}

func init() {
	//
}