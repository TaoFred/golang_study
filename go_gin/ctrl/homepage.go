package ctrl

import (
	"go_gin/common"
	"go_gin/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	SystemId int64
	entity.AddrInfo
}

// SearchAssetNumsAPI 首页资产统计
func SearchAssetNumsAPI(c *gin.Context) {
	var requestBody RequestBody
	if err := c.ShouldBind(&requestBody); err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_PARAM_ERROR, err.Error()))
		return
	}

	devLst, err := entity.GetAssetsNums(requestBody.SystemId, requestBody.AddrInfo)
	if err != nil {
		c.JSON(http.StatusOK, common.ResponseErr(common.RESULT_ERROR, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.ResponseOK(devLst))
}

func SearchScaleNumsAPI(c *gin.Context) {

}

func SearchCfgLogsAPI(c *gin.Context) {

}

func SearchCfgDefectsAPI(c *gin.Context) {

}

func SearchBackupListAPI(c *gin.Context) {

}
