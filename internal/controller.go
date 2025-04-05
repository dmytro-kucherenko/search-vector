package internal

import (
	"search-vector/internal/common/utils"
	"search-vector/internal/modules/document"
	"search-vector/internal/modules/query"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	documentService  *document.Service
	operationService *query.Service
}

func NewController(documentService *document.Service, operationService *query.Service) *Controller {
	return &Controller{documentService, operationService}
}

func (controller *Controller) Init(group *gin.RouterGroup) {
	group.POST("/documents", utils.Route(controller.CreateDocuments))
	group.GET("/search", utils.Route(controller.Search))
}

func (controller *Controller) CreateDocuments(params document.CreateRequest) (document.CreateResponse, error) {
	return controller.documentService.Create(params)
}

func (controller *Controller) Search(params query.SearchRequest) (query.SearchResponse, error) {
	return controller.operationService.Search(params)
}
