package document

import (
	"mime/multipart"
)

type Item struct {
	Name    string `json:"name"`
	Content string `json:"-"`
	Vector  Vector `json:"vector"`
}

type CompareItem struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type CreateRequest struct {
	Files []*multipart.FileHeader `form:"files" validate:"required,min=1"`
}

type CreateResponse struct {
	Documents []string `json:"documents"`
}
