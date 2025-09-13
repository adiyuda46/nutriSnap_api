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
	ReqGiziDetail struct {
		Label string `json:"Label"`
	}
	RespGiziDetail struct {
		Label        string `json:"Label"`
		Energi       string `json:"Energi"`
		Lemak        string `json:"Lemak"`
		VitA         string `json:"VitA"`
		VitB1        string `json:"VitB1"`
		VitB2        string `json:"VitB2"`
		VitB3        string `json:"VitB3"`
		VitC         string `json:"VitC"`
		Karbo        string `json:"Karbo"`
		Protein      string `json:"Protein"`
		SeratPangan  string `json:"SeratPangan"`
		Kalsium      string `json:"Kalsium"`
		Fosfor       string `json:"Fosfor"`
		Natrium      string `json:"Natrium"`
		Kalium       string `json:"Kalium"`
		Tembaga      string `json:"Tembaga"`
		Besi         string `json:"Besi"`
		Seng         string `json:"Seng"`
		BKarotene    string `json:"BKarotene"`
		KarotenTotal string `json:"KarotenTotal"`
		Air          string `json:"Air"`
		Abu          string `json:"Abu"`
	}
)
