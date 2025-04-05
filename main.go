package main

import (
	"search-vector/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	internal.Init(&router.RouterGroup)

	router.Run(":8000")
}
