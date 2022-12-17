package circuitbreaker

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	port = "8080"
)

var startTime = time.Now()

func server() {
	e := gin.Default()
	e.GET("/ping", func(c *gin.Context) {
		if time.Since(startTime) < 5*time.Second {
			c.String(http.StatusInternalServerError, "pong")
			return
		}
		c.String(http.StatusOK, "pong")
	})

	fmt.Printf("Starting server at port %s\n", port)
	e.Run(fmt.Sprintf(":%s", port))
}
