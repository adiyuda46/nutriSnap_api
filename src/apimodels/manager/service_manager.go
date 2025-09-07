package manager

import "api_model_cnn/src/apimodels/service"

type ServiceManager interface {
	NutriSnapService() service.NutriSnapService
}

type serviceManager struct {
	repo RepoManager
}

// service model
func (s *serviceManager) NutriSnapService() service.NutriSnapService  {
	return service.CreateNutriSnapServiceImpl(s.repo.Repository(),s.repo.HTTPRequest())
}


func CreateServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repo: repo,
	}
}
