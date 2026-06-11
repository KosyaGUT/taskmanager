package task

import (
	"github.com/google/uuid"
)

type Task struct {
	Title       string
	Description string
	Author      string
}

type User struct {
	id        uuid.UUID
	Login     string
	FirstName string
	LastName  string
	Password  string
}
