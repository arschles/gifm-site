package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// Screencast model struct
type Screencast struct {
	ID         uuid.UUID `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	EpisodeNum int       `json:"episode_num" db:"episode_num"`
	Title      string    `json:"title" db:"title"`
	Intro      string    `json:"intro" db:"intro"`
	Markdown   string    `json:"markdown" db:"markdown"`
}

// String is not required by pop and may be deleted
func (s Screencast) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Screencasts is not required by pop and may be deleted
type Screencasts []Screencast

// String is not required by pop and may be deleted
func (s Screencasts) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Screencast) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Screencast) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Screencast) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
