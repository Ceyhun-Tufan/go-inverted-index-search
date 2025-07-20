package main

import (
	"gosearch/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	core.InitDB()
	r := gin.Default()
	registerRoutes(r)
	r.Run(":8080")
}

func registerRoutes(r *gin.Engine) {
	group := r.Group("gosearch/")
	group.GET("/search", search)
	group.POST("/add", addIndex)

}

func search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter is required"})
		return
	}

	tokenized := core.Tokenize(&query)

	if len(tokenized) == 1 {

		result := core.Search(tokenized[0])

		if result == nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "Couldn't find an id."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": result})

	} else {
		result := core.SearchMulti(tokenized)

		if result == nil {
			c.JSON(http.StatusBadRequest, gin.H{"result": "Couldn't find an id (multi)."})
			return
		}

		c.JSON(http.StatusOK, gin.H{"result": result})
	}

}

func addIndex(c *gin.Context) {
	var req struct {
		ID    int    `json:"id" binding:"required"`
		Title string `json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tokenized []string = core.Tokenize(&req.Title)

	err := core.InsertIndex(&tokenized, req.ID)

	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Index added successfully",
		"data":    tokenized, // not needed
	})
}
