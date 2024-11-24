package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Deleted   bool      `json:"deleted"`
}
