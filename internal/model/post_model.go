package model

import (
	"time"

	"github.com/google/uuid"
)

type PostReponse struct {
	ID        *uuid.UUID    `json:"id,omitempty"`
	Title     string        `json:"title,omitempty"`
	Content   string        `json:"content,omitempty"`
	CreatedAt *time.Time    `json:"created_at,omitempty"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
	User      *UserResponse `json:"user,omitempty"`
}

type PostRequest struct {
	UserId  *uuid.UUID `json:"-" validate:"required,max=100"`
	Title   string     `json:"title,omitempty"`
	Content string     `json:"content,omitempty"`
}

type SearchPostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Page    int    `json:"page,omitempty" validate:"min=1"`
	Size    int    `json:"size,omitempty" validate:"min=1,max=100"`
}

type GetPostRequest struct {
	ID string `json:"id" validate:"required,max=100"`
}
