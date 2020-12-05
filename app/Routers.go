package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

type controller func(*gin.Engine)

func InitRouter() {
	r := gin.New()
	setupRouter(r)

	if err := r.Run(":8080"); err != nil {
		log.Panicf("Jieba API Service failed to start, err: %v\n", err)
	}
}

func setupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/word_divider", wordDividerHandler)
	}
}
