package document

import (
	"errors"
	"fmt"
	"mime/multipart"
	"search-vector/internal/common/schema"
)

type Service struct {
	documents []Item
	vectorIDF Vector
}

func NewService() *Service {
	return &Service{make([]Item, 0), make(Vector)}
}

func (service *Service) Get() (documents []Item, err error) {
	if len(service.documents) == 0 {
		err = errors.New("documents were not found")

		return
	}

	return service.documents, nil
}

func (service *Service) Compare(query string) (items []CompareItem, err error) {
	documents, err := service.Get()
	if err != nil {
		return
	}

	vector, err := GetVector(query, service.vectorIDF)
	if err != nil {
		return
	}

	fmt.Println(vector)
	for _, document := range documents {
		fmt.Println(document.Vector)
		value := CompareVectors(vector, document.Vector)
		items = append(items, CompareItem{
			Name:  document.Name,
			Value: value,
		})
	}

	return
}

func (service *Service) Create(params CreateRequest) (response CreateResponse, err error) {
	length := len(params.Files)
	service.documents = make([]Item, 0, length)

	for _, header := range params.Files {
		var file multipart.File
		file, err = header.Open()
		if err != nil {
			err = errors.New("failed to open document")

			return
		}

		defer file.Close()

		bytes := make([]byte, header.Size)
		_, err = file.Read(bytes)
		if err != nil {
			err = errors.New("failed to read document")

			return
		}

		content := string(bytes)
		ok := schema.ValidateTerms(content)
		if !ok {
			err = errors.New("document contains invalid terms")

			return
		}

		document := Item{
			Name:    header.Filename,
			Content: content,
		}

		service.documents = append(service.documents, document)
		response.Documents = append(response.Documents, document.Name)
	}

	service.Refresh()

	return
}

func (service *Service) Refresh() {
	contents := make([]string, 0, len(service.documents))
	for _, document := range service.documents {
		contents = append(contents, document.Content)
	}

	service.vectorIDF = GetVectorIDF(contents)
	for i, document := range service.documents {
		var err error
		document.Vector, err = GetVector(document.Content, service.vectorIDF)

		if err == nil {
			service.documents[i] = document
		}
	}
}
