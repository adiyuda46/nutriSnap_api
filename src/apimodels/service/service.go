package service

import "api_model_cnn/src/apimodels/repository"

type NutriSnapService interface {
	GetEmailService(id int) (string, error)
}

type NutriSnapServiceImpl struct {
	NutriSnap repository.Repository
}

// GetEmail implements NutriSnapService.
func (ns *NutriSnapServiceImpl) GetEmailService(id int) (string, error) {
	result ,err := ns.NutriSnap.GetEmailRepository(id) 
	if err != nil {
		return "error get email service",err
	}
	return  result,nil
}

func CreateNutriSnapServiceImpl(nutriSnap repository.Repository) NutriSnapService {
	return &NutriSnapServiceImpl{
		NutriSnap: nutriSnap,
	}
}
