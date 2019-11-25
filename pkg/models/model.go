package models

import (
	"time"
)

//Times *
type Times struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
