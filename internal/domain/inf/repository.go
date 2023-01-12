package inf

import (
	"go-web/internal/repository"
)

//

type RepoInterface interface {
	New() (*repository.PsqlRepository, error)
}

func init() {
	//
}