package controller

import (
	"api_model_cnn/src/apimodels/model"
	"api_model_cnn/src/apimodels/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (n *NutriSnapController) getId(c *gin.Context) {
	var input model.GetEmailById
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleError(c, 400, 400, "Error request data", err, "Error ShouldBindJSON")
		return
	}

	result, err := n.nutriSnap.GetEmailService(input.Userid)
	if err != nil {
		utils.HandleError(c, 404, 404, "Error Get data", err, "Error Get Email")
		return
	}

	// Handling Success
	utils.HandleSuccess(c, http.StatusOK, 200, "data di temukan", gin.H{
		"Email": result,
	}, "email: ", "ok ")
}
func (n *NutriSnapController) Predict(c *gin.Context)  {
	var input model.PredictModel 
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.HandleError(c, 400, 400, "Error request data", err, "Error ShouldBindJSON")
		return
	}
	

}
