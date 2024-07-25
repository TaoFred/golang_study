package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RESULT_OK = iota
	RESULT_ERROR
	RESULT_PARAM_ERROR
	RESULT_NO_AUTHORITY_ERROR
)

type Response struct {
	Status  bool        `json:"Status"`
	Message interface{} `json:"Message,omitempty"`
	Data    interface{} `json:"Data,omitempty"`
	Code    int         `json:"Code"`
}

func RespOK(data interface{}) Response {
	return Response{Code: RESULT_OK, Status: true, Data: data}
}

func RespErr(code int, message interface{}) Response {
	return Response{Code: code, Message: message, Status: false}
}

func GinRespList(c *gin.Context, count int64, list interface{}) {
	c.JSON(http.StatusOK, RespOK(map[string]interface{}{"TotalCount": count, "Result": list}))
}
