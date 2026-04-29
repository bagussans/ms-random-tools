package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Student struct {
	bun.BaseModel `bun:"table:students"`

	ID          int64     `bun:",pk,autoincrement"`
	UserID      int64     `bun:"user_id"`
	Name        string    `bun:"name"`
	CreatedDate time.Time `bun:"created_date"`
	UpdatedDate time.Time `bun:"updated_date"`
	Gender      *string   `bun:"gender"`
	ClassID     *int64    `bun:"class_id"`
	SchoolID    *int64    `bun:"school_id"`

	// User    *User           `bun:"rel:belongs-to,join:user_id=id"`
	// Groups  []*StudentGroup `bun:"rel:has-many,join:id=student_id"`
}

type StudentFull struct {
	bun.BaseModel `bun:"table:students"`

	ID          int64      `bun:"id"`
	UserID      int64      `bun:"user_id"`
	Name        string     `bun:"name"`
	Gender      *string    `bun:"gender"`
	CreatedDate time.Time  `bun:"created_date"`
	UpdatedDate *time.Time `bun:"updated_date"`
	// GroupID     *int64     `bun:"group_id"`
	// GroupName   *string    `bun:"group_name"`
	ClassID   *int64  `bun:"class_id"`
	ClassName *string `bun:"class_name"`
	// LessonID    *int64     `bun:"lesson_id"`
	// LessonName  *string    `bun:"lesson_name"`
	SchoolID   *int64  `bun:"school_id"`
	SchoolName *string `bun:"school_name"`

	GroupsAndLesson []map[string]interface{} `bun:"groups_and_lesson"`
}
