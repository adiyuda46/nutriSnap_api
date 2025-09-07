package thirdparty

import (
	"api_model_cnn/src/apimodels/utils"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type HTTPRequest interface {
	Request(path string, headers, params, querys map[string]string, body []byte, result interface{}) (*resty.Response, error)
}

type httpRequest struct {
	c *resty.Request
}

func CreateThirdpartyRequest(c *resty.Client) HTTPRequest {
	return &httpRequest{
		c: c.SetRetryCount(0).
			RemoveProxy().
			SetDebug(false).
			SetTimeout(75 * time.Second).
			R(),
	}
}

// Request implements Thirdparty.
func (t *httpRequest) Request(path string, headers, params, querys map[string]string, body []byte, result interface{}) (*resty.Response, error) {
	var response *resty.Response
	var err error
	start := time.Now()
	if headers == nil {
		response, err = t.c.
			SetHeader("Content-Type", utils.CONTENT_TYPE.JSON).
			SetPathParams(params).
			SetQueryParams(querys).
			SetBody(body).
			SetResult(&result).
			Post(path)
	} else {
		response, err = t.c.
			SetHeaders(headers).
			SetPathParams(params).
			SetQueryParams(querys).
			SetBody(body).
			SetResult(&result).
			Post(path)
	}
	duration := time.Since(start)
	utils.LogSuccess(fmt.Sprintf("%f", duration.Seconds()), "Duration: ")
	utils.LogSuccess("StatusCode: ", response.Status())
	utils.LogSuccess(response.Request.Header, "Request Header")
	utils.Log3rdParty(response.Request.Method, response.Request.URL, string(body), string(response.Body()))
	return response, err
}