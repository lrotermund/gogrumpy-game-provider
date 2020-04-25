package game

import "github.com/google/uuid"

type Game struct {
	ID              uuid.UUID `json:"id,omitempty"`
	PreviewImage    Image     `json:"previewImage,omitempty"`
	Name            string    `json:"name,omitempty"`
	Description     string    `json:"description,omitempty"`
	LongDescription string    `json:"longDescription,omitempty"`
}
