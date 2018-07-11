package v1controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTimeStamp Method
func (c *Controller) GetTimeStamp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"timestamp": time.Now(),
	})
}
