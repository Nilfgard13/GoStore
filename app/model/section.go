package model

import "time"

type Section struct {
	ID         string `gorm:"size:36;no null;uniqueIndex;primary_key"`
	Name       string `gorm:"size:100;"`
	Slug       string `gorm:"size:100;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []Category
}
