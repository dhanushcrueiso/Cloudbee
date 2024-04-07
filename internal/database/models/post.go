package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Post struct {
	Id              uuid.UUID `gorm:"primaryKey"`
	Title           string
	Content         string
	Author          string
	PublicationDate time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	Tags            pq.StringArray `gorm:"type:[]text"`
}
