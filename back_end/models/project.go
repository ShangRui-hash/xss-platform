package models

import "time"

//Project 项目
type Project struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Desc      string    `json:"desc" db:"description"`
	URLKey    string    `json:"url_key" db:"url_key"`
	UserID    int64     `json:"user_id" db:"user_id"`
	LootCount int       `json:"loot_count"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
