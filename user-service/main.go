package main

import (
	"github.com/gin-gonic/gin"
	"user-service/model"
)

/*Main function*/
func main() {
	r := gin.Default()

	r.Run("0.0.0.0:8080")
}
