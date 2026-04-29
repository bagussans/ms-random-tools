package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Group struct {
	bun.BaseModel `bun:"table:groups"`

	ID        int64     `bun:",pk,autoincrement"`
	LessonID  int64     `bun:"lesson_id"`
	Name      string    `bun:"name"`
	CreatedAt time.Time `bun:"created_date"`

	// Lesson   *Lesson          `bun:"rel:belongs-to,join:lesson_id=id"`
	// Students []*StudentGroup  `bun:"rel:has-many,join:id=group_id"`
}
