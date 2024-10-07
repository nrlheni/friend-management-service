package utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Path      string      `json:"path"`
	TimeStamp string      `json:"timestamp"`
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Errors    []Error     `json:"errors"`
}

type Error struct {
	Attribute string `json:"attribute,omitempty"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Exception error  `json:"-"`
}

func NewHttpResponse(request *http.Request, statusCode int, message string, result interface{}, errors []Error) *Response {
	if statusCode == 0 {
		statusCode = 200
	}
	path := fmt.Sprintf("%s://%s%s", getScheme(request), request.Host, request.RequestURI)
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	return &Response{
		Path:      path,
		TimeStamp: formattedTime,
		Status:    statusCode,
		Message:   message,
		Data:      result,
		Errors:    errors,
	}
}

func SuccessResponse(c *gin.Context, respCode int, message string, result interface{}) {
	if message == "" {
		message = "Success"
	}
	res := NewHttpResponse(c.Request, 200, message, result, nil)
	c.JSON(200, res)
	c.Abort()
	return
}

func getScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}

	return "http"
}