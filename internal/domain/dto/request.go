package dto

//
///
type RequestDto struct {
	//
	Code string `form:"code" binding:"required" json:"code"`
	OpenId string `uri:"openid" binding:"required" json:"openid"`
	EntId string `form:"entid" binding:"required" json:"entid"`
	Url string `form:"url" binding:"required" json:"url"`
	BusinessId string `uri:"businessid" binding:"required" json:"businessid"`
	Name string `form:"name" binding:"required" json:"name"`
	Num string `form:"num" binding:"required" json:"num"`
	Type string `form:"type" binding:"required" json:"type"`
	Description string `form:"description" binding:"required" json:"description"`
	Detail string `form:"detail" binding:"required" json:"detail"`
	// OrdId string `form:"orderid" binding:"required" json:"orderid"`
	// Repid string `form:"repid" binding:"required" json:"repid"`
	Id string `form:"id" binding:"required" json:"id"`
	Sig string `form:"sig" binding:"required" json:"sig"`
	Files []byte `form:"files" binding:"required" json:"files"`
	//
}

func init () {
	//
}