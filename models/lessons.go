package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Lesson struct {
	bun.BaseModel `bun:"table:lessons"`

	ID          int64     `bun:",pk,autoincrement"`
	ClassID     int64     `bun:"class_id"`
	Name        string    `bun:"name"`
	Desc        string    `bun:"desc"`
	CreatedDate time.Time `bun:"created_date"`

	// Class  *Class   `bun:"rel:belongs-to,join:class_id=id"`
	// Groups []*Group `bun:"rel:has-many,join:id=lesson_id"`
}
