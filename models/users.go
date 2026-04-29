package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID       int64  `bun:",pk,autoincrement"`
	Sub      string `bun:"sub"`
	Name     string `bun:"name"`
	Email    string `bun:"email"`
	Pict     string `bun:"pict"`
	Provider string `bun:"provider"`

	CreatedDate time.Time `bun:"created_date"`
	UpdatedDate time.Time `bun:"updated_date"`

	// Schools []*School `bun:"rel:has-many,join:id=user_id"`
}
