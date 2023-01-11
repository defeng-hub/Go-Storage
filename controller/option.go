package controller

import (
	"encoding/json"
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOptions(c *gin.Context) {
	var options []*model.Option
	for k, v := range common.OptionMap {
		options = append(options, &model.Option{
			Key:   k,
			Value: common.Interface2String(v),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    options,
	})
	return
}

func GetNotice(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    common.OptionMap["Notice"],
	})
	return
}

func UpdateOption(c *gin.Context) {
	var option model.Option
	err := json.NewDecoder(c.Request.Body).Decode(&option)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	err = model.UpdateOption(option.Key, option.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data": gin.H{
			"version": common.Version,
		},
	})
	return
}
