package dto

//
///
type ConfigDto struct {
	DbConnStr string `json:"dbconnstr"`
	Auth_code2Session string `json:"auth_code2session"`
	Etcd string `json:"etcd"`
}

func init () {
	//
}