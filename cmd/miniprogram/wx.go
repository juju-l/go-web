package main

import (
	"time"
	"encoding/base64"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"go-web/internal/domain/inf"
	"go-web/internal/service"
	"go-web/internal/domain/dto"
	"go-web/util/ext"
	"os"
	"io/ioutil"
)

var svc inf.ApiInterface

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.Use(ext.Authorize())
	reg();r.Use(c())//#
			grp := func (s string) {
	api := r.Group(s); {
			api.GET("/auth", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.Auth( req )) }, /**/)//
			api.POST("/upContents", func(c *gin.Context) { var req=&dto.RequestDto{}
			c.ShouldBind(req);if c.Request.Body!=nil{req.Files,_=ioutil.ReadAll(c.Request.Body); req.Sig=sig(); os.WriteFile(req.Sig,req.Files,0600)};c.ShouldBindUri(req)
			c.JSON(200, svc.UpContents( req )) }, /**/)//
			api.GET("/getAvatarUrl/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetAvatarUrl( req )) }, /**/)//
			api.GET("/getBusinessInfo/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetBusinessInfo( req )) }, /**/)//
			api.GET("/upBusinessInfo/:openid/:businessid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.UpBusinessInfo( req )) }, /**/)//
			api.GET("/getDetails/:openid/:type", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetDetails( req )) }, /**/)//
			api.GET("/getOrder/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetOrder( req )) }, /**/)//
			api.POST("/getCnts", func(c *gin.Context) { var req=&dto.RequestDto{}
			c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req)
			c.JSON(200, svc.GetCnts( req )) }, /**/)//
			api.GET("/getRepair/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetRepair( req )) }, /**/)//
	}
			}
	grp("/api");grp("/v1");grp("/v2")
	r.Use(cors.Default())
	//r.Use(c())
	r.RunTLS(":443", "config/tls.crt", "config/.tls.key")
}

//

func sig() string {
	return base64.StdEncoding.EncodeToString(
	[]byte(
	strconv.FormatInt(time.Now().UnixNano(),10),
	),
	)
}

func init() {
	//
	svc = 
			&service. ///
	ApiService{}
	//
}