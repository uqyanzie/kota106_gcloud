package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uqyanzie/kota106_gcloud/model"
	"github.com/uqyanzie/kota106_gcloud/table"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK,
			gin.H{
				"message": "Here is the available table",
				"table":   table.GetTableInfo(),
			},
		)
	})

	r.POST("/query", func(c *gin.Context) {
		var request model.Request

		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"generated_sql": request.GetProjection(),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
