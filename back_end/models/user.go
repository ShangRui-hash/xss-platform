package models

import "time"

//User 用户
type User struct {
	ID        int64     `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password,omitempty"`
	LoginedAt string    `json:"logined_at"`
	IsBanned  bool      `json:"is_banned"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
