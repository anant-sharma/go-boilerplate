package clock

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
GetTime function to authenticate user
*/
func GetTime(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"timestamp": time.Now(),
	})
}
