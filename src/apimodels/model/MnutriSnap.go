package model

type (
	GetEmailById struct {
		Userid int `json:"Userid"`
	}
	PredictModel struct {
		PredictModelImg string `json:"PredictModelImg"`
	}

	ReqPredict struct {
		File string `json:"file"`
	}
	ResponsePredict struct {
		Confidence float64 `json:"confidence"`
		Label string `json:"label"`
	}
)