package main

import (
	// "github.com/..."
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"go-web/internal/domain/inf"
	"go-web/internal/service"
	"go-web/internal/domain/dto"
	"go-web/util/ext"
)

var svc inf.ApiInterface

func main() {
	r := gin.Default()
	r.Use(ext.Authorize())
	//#reg();r.Use(c())
	api := r.Group("/api"); {
			api.GET("/auth", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.Auth( req )) }, /**/)//
			api.GET("/getAvatarUrl/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetAvatarUrl( req )) }, /**/)//
			api.GET("/getBusinessInfo/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetBusinessInfo( req )) }, /**/)//
			api.GET("/upBusinessInfo/:openid/:businessid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.UpBusinessInfo( req )) }, /**/)//
			api.GET("/getDetails/:openid/:type", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetDetails( req )) }, /**/)//
			api.GET("/getOrder/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetOrder( req )) }, /**/)//
			api.GET("/getRepair/:openid", func(c *gin.Context) { var req=&dto.RequestDto{}; c.ShouldBind(req);c.ShouldBindJSON(req);c.ShouldBindUri(req); c.JSON(200, svc.GetRepair( req )) }, /**/)//
	}
	r.Use(cors.Default())
	//r.Use(c())
	r.RunTLS(":443", "config/tls.crt", "config/.tls.key")
}

func init() {
	//
	svc = 
			&service. ///
	ApiService{}
	//
}