package middleware

import (
	"github.com/defeng-hub/Go-Storage/common"
	"github.com/defeng-hub/Go-Storage/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 检查是否有权限
func permissionCheckHelper(c *gin.Context, requiredPermission int) {
	c.Set("role", common.RoleGuestUser)
	session := sessions.Default(c)
	role := session.Get("role")
	username := session.Get("username")
	if username != nil {
		c.Set("username", username)
	} else {
		// Check token
		token := c.Request.Header.Get("Authorization")
		user := model.ValidateUserToken(token)
		if user != nil && user.Username != "" {
			// Token is valid
			username = user.Username
			role = user.Role
			c.Set("username", username)
		} else {
			c.Set("username", "访客用户")
		}
	}
	if requiredPermission == common.RoleGuestUser {
		c.Next()
		return
	}
	if role == nil || role.(int) < requiredPermission {
		if c.Request.URL.Path == "/" {
			c.HTML(http.StatusMovedPermanently, "login.html", gin.H{
				"message": "无权访问此页面，请检查你是否登录或者是否有相关权限",
			})

		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "无权进行此操作，请检查你是否登录或者是否有相关权限",
			})
		}
		c.Abort()
		return
	}
	c.Set("role", role)
	c.Next()
}

func ImageUploadPermissionCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		atoi, err := strconv.Atoi(common.OptionMap["ImageUploadPermission"])
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "服务器出岔子了，err:" + err.Error(),
			})
			c.Abort()
			return
		}
		permissionCheckHelper(c, atoi)
	}
}

func FileUploadPermissionCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		atoi, err := strconv.Atoi(common.OptionMap["FileUploadPermission"])
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "服务器出岔子了，err:" + err.Error(),
			})
			c.Abort()
			return
		}
		permissionCheckHelper(c, atoi)
	}
}

// ShowIndexPermissionCheck 是否可以显示首页
func ShowIndexPermissionCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		i, err := strconv.Atoi(common.OptionMap["IndexPermission"])
		if err != nil {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"message": "服务器出岔子了，err:" + err.Error(),
				"option":  common.OptionMap,
			})
			c.Abort()
			return
		}
		permissionCheckHelper(c, i)
	}
}
