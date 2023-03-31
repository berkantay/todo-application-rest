package model

import (
	"github.com/google/uuid"
)

type User struct {
	UUID     uuid.UUID `json:"uuid,omitempty"`
	Username *string   `json:"username,omitempty"`
	Password *string   `json:"Password,omitempty"`
}
