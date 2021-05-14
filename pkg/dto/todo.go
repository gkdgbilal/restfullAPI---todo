package dto

import (
	"RestFullAPI-todo/pkg/enums"
	"time"
)

type TodoDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title" validate:"required,max=21"`
	Description string `json:"description" validate:"required,max=21"`
	Completed   bool   `json:"completed"`
}
type TodoResponse struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Completed   bool         `json:"completed"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Status      enums.Status `json:"status"`
}
