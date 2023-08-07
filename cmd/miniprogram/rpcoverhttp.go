package main

import (
	// "github.com/..."
	"github.com/gin-gonic/gin"
	"strings"
)

func c() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ProtoMajor == 2 && strings.HasPrefix(c.GetHeader("Content-Type"),"application/grpc") && true {
			c.Status(200)
			s.ServeHTTP(c.Writer,/*->,*/c.Request)
			c.Abort();/*must->;*/return
		}; c.Next(); //
	}
}

//

func init() {
	//
}