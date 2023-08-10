package controller

import (
	"github.com/gin-gonic/gin"
	"thingue-launcher/common/model"
)

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}) {
	json := &model.JsonStruct{Code: code, Msg: msg, Data: data}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &model.JsonStruct{Code: code, Msg: msg}
	c.JSON(200, json)
}
