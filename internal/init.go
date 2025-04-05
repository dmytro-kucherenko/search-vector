package internal

import (
	"search-vector/internal/modules/document"
	"search-vector/internal/modules/query"

	"github.com/gin-gonic/gin"
)

func Init(group *gin.RouterGroup) error {
	documentService := document.NewService()
	queryService := query.NewService(documentService)

	controller := NewController(documentService, queryService)
	controller.Init(group)

	return nil
}
