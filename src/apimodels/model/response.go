package model

// ResponseWrapper is a struct that wrap response of the system
type (
	ResponseWrapper struct {
		Message  string      `json:"message"`
		Response int         `json:"response"`
		Result   interface{} `json:"result"`
	}
)
