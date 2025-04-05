package query

import "search-vector/internal/modules/document"

type OprationType int

const (
	OperationAnd OprationType = iota
	OperationOr  OprationType = iota
)

type Operation struct {
	Type       OprationType
	Terms      []string
	Operations []Operation
}

type SearchRequest struct {
	Query string `form:"query" validate:"required,lowercase,query"`
}

type SearchResponse struct {
	Documents []document.CompareItem `json:"documents"`
}
