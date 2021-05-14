package dto

type ListResponse struct {
	Page          int64       `json:"page"`
	Size          int64       `json:"size"`
	TotalPages    int64       `json:"totalPages"`
	TotalElements int64       `json:"totalElements"`
	IsLastPage    bool        `json:"isLastPage"`
	HasNext       bool        `json:"hasNext"`
	Content       interface{} `json:"content"`
}
