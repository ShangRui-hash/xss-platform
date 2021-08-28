package models

import "time"

//Module 模块
type Module struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	Desc       string    `db:"description" json:"desc"`
	XSSPayload string    `db:"xss_payload" json:"xss_payload,omitempty"`
	UserID     int64     `db:"user_id" json:"user_id,omitempty"`
	IsAdmin    *bool     `db:"user_type" json:"is_admin,omitempty"`
	Username   string    `db:"username" json:"username,omitempty"`
	IsCommon   *bool     `db:"is_common" json:"is_common,omitempty"`
	ParamList  []string  `json:"param_list"`
	OptionList []string  `json:"option_list"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}
