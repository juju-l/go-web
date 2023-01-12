package dto

//
///
type ResponseDto struct {
	ErrCode int `json:"errCode"`
	Data map[string]any `json:"data"` //
	ErrMsg string `json:"errMsg"`
}

func init () {
	//
}