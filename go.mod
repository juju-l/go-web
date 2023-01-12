module go-web

//replace (
// openapi =>
//	codeup.aliyun.com/vipex/go-v3openapi
//v0.1.0
//)
replace openapi => codeup.aliyun.com/vipex/go-v3openapi v0.1.0
require (
	//
	gorm.io/gorm v1.24.1
	github.com/gin-gonic/gin v1.8.1
	github.com/golang-jwt/jwt/v4 v4.4.3 // Authz
	gorm.io/driver/postgres v1.4.5
	github.com/gin-contrib/cors v1.4.0
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	gopkg.in/yaml.v2 v2.4.0
	//
)
require openapi v0.0.0-00010101000000-000000000000
exclude (
	//
)

go 1.19