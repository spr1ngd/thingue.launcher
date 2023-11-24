package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageResult[T any] struct {
	List     []T   `json:"list"`
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

type Response[T any] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 200
)

func Result[T any](code int, data T, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response[T]{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData[T any](data T, c *gin.Context) {
	Result(SUCCESS, data, "查询成功", c)
}

func OkWithDetailed[T any](data T, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed[T any](data T, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
