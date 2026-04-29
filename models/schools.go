package models

import (
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

type School struct {
	bun.BaseModel `bun:"table:schools"`

	ID          int64     `bun:",pk,autoincrement"`
	UserID      int64     `bun:"user_id"`
	Name        string    `bun:"name"`
	Location    string    `bun:"location"`
	CreatedDate time.Time `bun:"created_date"`
	UpdatedDate time.Time `bun:"updated_date"`
	IsDeleted   bool      `bun:"is_deleted"`

	// User   *User    `bun:"rel:belongs-to,join:user_id=id"`
	// Classes []*Class `bun:"rel:has-many,join:id=school_id"`
}

type SchoolWithClassRegistered struct {
	bun.BaseModel `bun:"table:schools"`

	ID          int64     `bun:",pk,autoincrement"`
	UserID      int64     `bun:"user_id"`
	Name        string    `bun:"name"`
	Location    string    `bun:"location"`
	CreatedDate time.Time `bun:"created_date"`
	UpdatedDate time.Time `bun:"updated_date"`
	IsDeleted   bool      `bun:"is_deleted"`

	//param for join
	Classes json.RawMessage `bun:"classes"`

	// User   *User    `bun:"rel:belongs-to,join:user_id=id"`
	// Classes []*Class `bun:"rel:has-many,join:id=school_id"`
}
