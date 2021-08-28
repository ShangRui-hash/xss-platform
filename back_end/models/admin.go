package models

import "time"

//Admin 管理员
type Admin struct {
	ID        int64     `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
