package utils

import (
	modelapp "api_model_cnn/src/apimodels/model"
	"encoding/json"

	"github.com/alexcesaro/log/stdlog"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

var Logkoe = stdlog.GetFromFlags()
var CONTENT_TYPE = modelapp.ContentType{
	JSON:     "application/json",
	FormData: "multipart/form-data; boundary=<calculated when request is sent>",
	Plain:    "text/plain",
}


// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// LogSuccess is a function of success process that only generate log to console
func LogSuccess(detail interface{}, info string) {
	// // Check if data has base64, it's needed to keep log clean from long base64
	// detail = IsContainsBase64(detail)

	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
	}).Info(info)
	Logkoe.Info("msg =", info, "detail =", detail)
}

// LogError is a function of failed process that only generate log to console
func LogError(detail interface{}, info string) {
	// // Check if data has base64, it's needed to keep log clean from long base64
	// detail = IsContainsBase64(detail)

	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
	}).Error(info)
	Logkoe.Error("msg =", info, "detail =", detail)
}
// HandleError is a function of failed process that send to the front as a response error
func HandleError(c *gin.Context, status int, code int, message string, detail interface{}, info string) {
	// Assign struct ResponseWrapper with parameter above
	response := modelapp.ResponseWrapper{
		Message:  message,
		Response: code,
		Result:   nil,
	}

	// Write header manually
	c.Header("Content-Type", "application/json")

	// Use JSON as a response and send initialize struct above
	c.JSON(status, response)

	//Get reqid
	reqid := requestid.Get(c)

	// // Check if data has base64, it's needed to keep log clean from long base64
	// detail = IsContainsBase64(detail)

	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
		"Id":     reqid,
	}).Error(info)
	Logkoe.Error(
		"Resp => method =", c.Request.Method,
		"||Id =", reqid,
		"||url =", c.Request.URL,
		"||msg =", info,
		"||detail =", detail)
}


// HandleSuccess is a function of success process that send to the front as a response success
func HandleSuccess(c *gin.Context, status int, code int, message string, data map[string]interface{}, detail interface{}, info string) {
	// Assign struct ResponseWrapper with parameter above
	response := modelapp.ResponseWrapper{
		Message:  message,
		Response: code,
		Result:   data,
	}

	// Write header manually
	c.Header("Content-Type", "application/json")

	// Use JSON as a response and send initialize struct above
	c.JSON(status, response)

	// Get reqid
	reqid := requestid.Get(c)

	// // Check if data has base64, it's needed to keep log clean from long base64
	// detail = IsContainsBase64(detail)

	// Assign detail with parameter above and generate console
	logger.WithFields(logger.Fields{
		"detail": detail,
		"Id":     reqid,
	}).Info(info)
	Logkoe.Info(
		"Resp => method =", c.Request.Method,
		"||Id =", reqid,
		"||url =", c.Request.URL,
		"||msg =", info,
		"||detail =", detail)
}

// Log3rdParty is a function to create format log contains method, url, and request, response body. It's created after consume 3rd party.
func Log3rdParty(method, url, req, resp interface{}) {

	// Assign detail with parameter above and generate console
	// Request
	logger.WithFields(logger.Fields{
		"method":  method,
		"url":     url,
		"request": req,
	}).Info("Resty Request")
	Logkoe.Info("Resty Request => method =", method,
		"||url =", url,
		"||request =", req)

	// Response
	logger.WithFields(logger.Fields{
		"response": resp,
	}).Info("Resty Response")
	Logkoe.Info("Resty Response ==> response =", resp)
}

// ConvertResponse is a function to change struct to map interface like response format
func ConvertResponse(data interface{}) (map[string]interface{}, string) {
	var newData map[string]interface{}
	dt, _ := json.Marshal(data)
	json.Unmarshal(dt, &newData)

	return newData, string(dt)
}