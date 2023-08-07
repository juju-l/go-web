package token

//
///
type TknInterface interface {
	New(...string) (interface{}, error)
	ValidateTkn(token string) (bool, error)
	GenToken() (string, error)
}

func init() {
	//
}