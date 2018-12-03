package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qdk/qdk.v1/service"
)

type (
	UltraSonicInput struct {
		Distance string `json:"distance" binding:"required"`
	}
)

func (s *Server) PostUltraSonicHandler(c *gin.Context) {
	deviceId := c.MustGet("deviceId").(string)
	input := &UltraSonicInput{}
	if err := c.ShouldBind(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		s.PiMgr.PostUltraSonic(
			c.Request.Context(),
			service.PostUltraSonicInput{
				Distance: input.Distance,
				DeviceId: deviceId,
			})
		c.JSON(200, gin.H{"message": "ok"})
	}
}
