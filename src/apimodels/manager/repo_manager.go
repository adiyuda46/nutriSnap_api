package manager

import (
	"api_model_cnn/src/apimodels/connection"
	"api_model_cnn/src/apimodels/repository"
	"api_model_cnn/src/apimodels/thirdparty"

	"github.com/go-resty/resty/v2"
)

type RepoManager interface {
	Repository() repository.Repository
	HTTPRequest() thirdparty.HTTPRequest
}

type repoManager struct {
	resty    *resty.Client
	postgres connection.Connection
}

// repository model
func (r *repoManager) Repository() repository.Repository {
	return repository.CreateRepository(r.postgres.SqlDb())
}

// HTTPRequest
func (r *repoManager) HTTPRequest() thirdparty.HTTPRequest {
	return thirdparty.CreateThirdpartyRequest(r.resty)
}

func CreateRepoManager(postgres connection.Connection) RepoManager {
	return &repoManager{
		resty:    resty.New(),
		postgres: postgres,
	}
}
