package ctrl

import (
	"go_gin/common"
	"go_gin/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddSystem(c *gin.Context) {
	var reqBody entity.SystemRequest
	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_PARAM_ERROR, err.Error()))
		return
	}

	systemInfo, err := entity.AddSystem(&reqBody)
	if err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_ERROR, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.ResponseOK(map[string]interface{}{"SystemId": systemInfo.SystemId}))
}

func DeleteSystem(c *gin.Context) {
	var reqBody struct {
		SystemIds string
	}
	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_PARAM_ERROR, err.Error()))
		return
	}

	entity.DeleteSystem(reqBody.SystemIds)
}

func ModifySystem(c *gin.Context) {
	var reqBody entity.SystemRequest
	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_PARAM_ERROR, err.Error()))
		return
	}

	err := entity.ModifySystem(&reqBody)
	if err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_ERROR, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.ResponseOK(0))
}

func GetAllSystem(c *gin.Context) {
	var reqBody entity.SystemRequest
	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_PARAM_ERROR, err.Error()))
		return
	}

}

func GetSystem(c *gin.Context) {
	deviceIdStr := c.Query("DeviceId")
	deviceId, _ := strconv.ParseInt(deviceIdStr, 10, 64)
	systemList, err := entity.GetSystemByDevId(deviceId)
	if err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_ERROR, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.ResponseOK(systemList))
}
