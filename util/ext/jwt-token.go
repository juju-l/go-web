package ext

import (
	"time"
	"github.com/golang-jwt/jwt/v4"
	//"fmt"
)

var wxSigningKey = []byte("wx123789456cc") //

func New(
			entid string, is_authorization bool, openid string,
	) (*WxClaims, error) {
	return &WxClaims{entid,is_authorization,jwt.RegisteredClaims {
			ID:openid, ExpiresAt:jwt.NewNumericDate(time.Now().Add(7200 * time.Second)), Issuer:"wxVipex.cc",
		}}, nil
}

func ValidateTkn(token string) (
			bool, error,
	) {
	v := false
	if t,e := jwt.Parse(token,
		func(tkn *jwt.Token) (interface{},error) {
			return wxSigningKey, nil
		},
	/**/);e != nil {return false, e} else {
	v = t.Valid
	}
	return v, nil
}

///
// func
//***

func GenToken(w *WxClaims) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, w).SignedString(wxSigningKey)
}

type WxClaims struct {
	EntId string `json:"entid"` //
	IsAuthorization bool `json:"is_authorization"` //
	jwt.RegisteredClaims
}

func init() {
	//
}