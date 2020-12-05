package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanyiwu/gojieba"
)

func wordDividerHandler(c *gin.Context) {
	var words []string
	sentence := c.Query("sentence")
	if sentence == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing params[:sentence]",
		})
		return
	}

	parser := gojieba.NewJieba()
	defer parser.Free()

	words = parser.Cut(sentence, true)
	c.JSON(http.StatusOK, gin.H{
		"result": words,
	})
}
