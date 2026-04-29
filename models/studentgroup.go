package models

import (
	"time"

	"github.com/uptrace/bun"
)

type StudentGroup struct {
	bun.BaseModel `bun:"table:student_groups"`

	ID        int64     `bun:",pk,autoincrement"`
	GroupID   int64     `bun:"group_id"`
	StudentID int64     `bun:"student_id"`
	CreatedAt time.Time `bun:"created_date"`
	UpdatedAt time.Time `bun:"updated_date"`

	// Group   *Group   `bun:"rel:belongs-to,join:group_id=id"`
	// Student *Student `bun:"rel:belongs-to,join:student_id=id"`
}
