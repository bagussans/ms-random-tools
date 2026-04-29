package models

import (
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

type ClassModel struct {
	bun.BaseModel `bun:"table:classes"`

	ID        int64     `bun:",pk,autoincrement"`
	SchoolID  int64     `bun:"school_id"`
	Name      string    `bun:"name"`
	CreatedAt time.Time `bun:"created_date"`
	IsDeleted bool      `bun:"is_deleted"`
	Year      int64     `bun:"year"`

	// School *School   `bun:"rel:belongs-to,join:school_id=id"`
	// Lessons []*Lesson `bun:"rel:has-many,join:id=class_id"`
}

type ClassWithLesson struct {
	bun.BaseModel `bun:"table:schools"`

	ID        int64     `bun:",pk,autoincrement"`
	SchoolID  int64     `bun:"school_id"`
	Name      string    `bun:"name"`
	CreatedAt time.Time `bun:"created_date"`
	IsDeleted bool      `bun:"is_deleted"`
	Year      int64     `bun:"year"`

	//param for join
	Lessons json.RawMessage `bun:"lessons"`

	StudentCount int64 `bun:"student_count"`
	GroupCount   int64 `bun:"group_count"`

	// User   *User    `bun:"rel:belongs-to,join:user_id=id"`
	// Classes []*Class `bun:"rel:has-many,join:id=school_id"`
}
