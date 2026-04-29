package models

import (
	"time"

	"github.com/uptrace/bun"
)

type StudentLesson struct {
	bun.BaseModel `bun:"table:student_lessons"`

	ID           int64     `bun:",pk,autoincrement"`
	StudentID    int64     `bun:"student_id"`
	LessonID     int64     `bun:"lesson_id"`
	PreGroupDest int64     `bun:"pre_group_dest"`
	CreatedDate  time.Time `bun:"created_date"`

	// Lesson  *Lesson  `bun:"rel:belongs-to,join:lesson_id=id"`
	// Student *Student `bun:"rel:belongs-to,join:student_id=id"`
}
