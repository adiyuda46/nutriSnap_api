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
		Label      string  `json:"label"`
	}
	RespGizi struct {
		Label         string `json:"Label"`
		Energi        string `json:"Energi"`
		Protein       string `json:"Protein"`
		Lemak         string `json:"Lemak"`
		Karbo         string `json:"Karbo"`
		GiziUnggulan1 string `json:"GiziUnggulan1"`
		GiziUnggulan2 string `json:"GiziUnggulan2"`
		GiziUnggulan3 string `json:"GiziUnggulan3"`
	}
)
