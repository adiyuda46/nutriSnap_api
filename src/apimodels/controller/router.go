package controller

import (
	"api_model_cnn/src/apimodels/service"

	"github.com/gin-gonic/gin"
)

type NutriSnapController struct {
	r         *gin.Engine
	nutriSnap service.NutriSnapService
}

func CreateNutriSnapController(r *gin.Engine, nutriSnap service.NutriSnapService) {
	controller := NutriSnapController {
		r : r,
		nutriSnap: nutriSnap,
	}

	//public
	V1public := r.Group("/api/v1/bafpayment/public")
	{
		V1public.POST("/", controller.r.HandleContext)
		V1public.POST("/id",controller.getId)
		//V1public.POST("/predict")
	}
}
