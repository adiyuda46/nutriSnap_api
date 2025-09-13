package service

import (
	"api_model_cnn/src/apimodels/model"
	"api_model_cnn/src/apimodels/repository"
	"api_model_cnn/src/apimodels/thirdparty"
	"api_model_cnn/src/apimodels/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

type NutriSnapService interface {
	GetEmailService(id int) (string, error)
	PredictService(string) (model.RespGizi, error)
}

type NutriSnapServiceImpl struct {
	NutriSnap repository.Repository
	Thirdparty thirdparty.HTTPRequest 
}

// PredictService implements NutriSnapService.
func (ns *NutriSnapServiceImpl) PredictService(imgBase64 string) (model.RespGizi, error) {
	headers := map[string]string{
		"Content-Type": utils.CONTENT_TYPE.JSON,
	}

	req := model.ReqPredict{
		File: imgBase64,
	}
	reqBody, err := json.Marshal(req)
	if err != nil {
		utils.LogError(err, "Error marshalling request body")
		return model.RespGizi{}, fmt.Errorf("Error marshalling request body")
	}
	var resp model.ResponsePredict
	res, err := ns.Thirdparty.Request(viper.GetString("url.predict"), headers, nil, nil, reqBody, &resp)
	if err != nil {
		utils.LogError(err, "Error in Consume 3rd Party")
		return model.RespGizi{}, fmt.Errorf("Error in Consume 3rd Party")
	}

	if err := json.Unmarshal(res.Body(), &resp); err != nil {
		utils.LogError(err, "Error unmarshalling response body")
		return model.RespGizi{}, fmt.Errorf("Error unmarshalling response body")
	}

	// get data gizi
	result,err := ns.NutriSnap.GetGizi(resp.Label)
	if err != nil {
		return model.RespGizi{}, fmt.Errorf("data not found")
	}


	return result , nil
}

// GetEmail implements NutriSnapService.
func (ns *NutriSnapServiceImpl) GetEmailService(id int) (string, error) {
	result, err := ns.NutriSnap.GetEmailRepository(id)
	if err != nil {
		return "error get email service", err
	}
	return result, nil
}

func CreateNutriSnapServiceImpl(nutriSnap repository.Repository,  thirdpartyClient thirdparty.HTTPRequest) NutriSnapService {
	return &NutriSnapServiceImpl{
		NutriSnap: nutriSnap,
		Thirdparty: thirdpartyClient,
	}
}
