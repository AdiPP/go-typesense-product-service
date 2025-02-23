package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type pingHandler struct {
}

func (h *pingHandler) handle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func newPingHandler() *pingHandler {
	o := new(pingHandler)
	return o
}
