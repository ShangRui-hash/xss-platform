package controllers

import "xss/models"

type _RespAdminLogin struct {
	Code  ResCode     `json:"code"`
	Msg   interface{} `json:"msg"`
	Token string      `json:"token"` //jwt令牌
}

type _RespProjectForm struct {
	Code ResCode                   `json:"code"`
	Msg  interface{}               `json:"msg"`
	Form models.ParamCreateProject `json:"form"`
}

type _RespModuleDetail struct {
	Code ResCode       `json:"code"`
	Msg  interface{}   `json:"msg"`
	Form models.Module `json:"module"`
}
