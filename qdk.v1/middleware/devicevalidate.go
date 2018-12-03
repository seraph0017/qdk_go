package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qdk/qdk.v1/service"
)

func DeviceValidate(pi service.PiMgr) gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceId := c.GetHeader("X-device")
		if ok := pi.Validate(c.Request.Context(), deviceId); ok {
			c.Set("deviceId", deviceId)
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "device not found"})
			c.Abort()
		}
	}
}
