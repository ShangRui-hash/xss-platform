package models

import "time"

//Loot 战利品
type Loot struct {
	ID        int64     `json:"id" db:"id"`
	URLKey    string    `json:"url_key" db:"url_key"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
