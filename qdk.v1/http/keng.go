package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) KengIndexHandler(c *gin.Context) {
	resp, err := s.KengMgr.GetAllKeng()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
