package query

import "search-vector/internal/modules/document"

const Threshold = 0.9

type Service struct {
	documentService *document.Service
}

func NewService(documentService *document.Service) *Service {
	return &Service{documentService}
}

func (service *Service) Search(params SearchRequest) (response SearchResponse, err error) {
	documents, err := service.documentService.Compare(params.Query)
	if err != nil {
		return
	}

	matchedDocuments := make([]document.CompareItem, 0, len(documents))
	for _, document := range documents {
		if document.Value > Threshold {
			matchedDocuments = append(matchedDocuments, document)
		}
	}

	response.Documents = matchedDocuments

	return
}
