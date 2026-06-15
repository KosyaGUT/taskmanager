package task

import (
	"github.com/google/uuid"
)

type Task struct {
	Id           uuid.UUID `json:"id"`
	CreatorLogin string    `json:"creator_login"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	Author       string    `json:"author,omitempty"`
}

type User struct {
	Id        uuid.UUID `json:"id"`
	Login     string    `json:"login"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
}
