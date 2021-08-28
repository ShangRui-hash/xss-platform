package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//ValidateJSONParam 效验JSON传参，如果参数不符合效验规则并返回响应
func ValidateJSONParam(c *gin.Context, param interface{}) error {
	err := c.ShouldBindJSON(param)
	if nil == err {
		return nil
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		//非validator.ValidationErrors类型错误
		RespErr(c, CodeInvalidParam)
		return err
	}
	//validator.ValidationErrors类型错误则进行翻译
	RespErrMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
	return errs
}
