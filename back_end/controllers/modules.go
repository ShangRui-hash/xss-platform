package controllers

import (
	"fmt"
	"strconv"
	"xss/dao/mysql"
	"xss/logic"
	"xss/models"
	"xss/pkg/validate"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// AddModuleHandler 添加xss模块
// @Summary 添加xss模块
// @Description 添加xss模块
// @Tags 前后台共用接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 令牌"
// @Param object query models.ParamsAddModule true "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} Resp
// @Router /api/v1/modules [post]
// @Router /api/v1/admin/modules [post]
func AddModuleHandler(c *gin.Context) {
	//1.接受参数
	var params models.ParamsAddModule
	if err := ValidateJSONParam(c, &params); err != nil {
		zap.L().Error("add module with invalid params", zap.Error(err))
		return
	}
	userid, exist := c.Get(CtxUserIDKey)
	if !exist {
		zap.L().Error("userid don't exist")
		RespErr(c, CodeServerBusy)
		return
	}
	username, exist := c.Get(CtxUsernameKey)
	if !exist {
		zap.L().Error("username don't exist")
		RespErr(c, CodeServerBusy)
		return
	}
	isAdmin, exist := c.Get(CtxIsAdminKey)
	if !exist {
		zap.L().Error("c.Get(CtxIsAdminKey) failed")
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑
	m, err := logic.AddModule(params, userid.(int64), isAdmin.(bool))
	if err != nil {
		zap.L().Error("logic.AddModule failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	m.Username = username.(string)
	//3.返回响应
	RespSuc(c, m)
}

//GetModulesHandler 获取模块
func GetModulesHandler(c *gin.Context) {
	//1.接收参数,后端效验
	offset, err := strconv.ParseInt(c.Param("offset"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidOffset)
		return
	}
	count, err := strconv.ParseInt(c.Param("count"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidCount)
		return
	}
	//2.业务逻辑
	modules, err := logic.GetModules(offset, count)
	if err != nil {
		zap.L().Error("logic.GetModules failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, modules)
}

//GetCommonModulesHandler 获取公共模块
func GetCommonModulesHandler(c *gin.Context) {
	//1.接收参数,后端效验
	offset, err := strconv.ParseInt(c.Param("offset"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidOffset)
		return
	}
	count, err := strconv.ParseInt(c.Param("count"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidCount)
		return
	}
	//2.业务逻辑
	modules, err := logic.GetCommonModules(offset, count)
	if err != nil {
		zap.L().Error("logic.GetCommonModules failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, modules)
}

//AdminGetModuleDetailHandler 管理员获取模块详情
func AdminGetModuleDetailHandler(c *gin.Context) {
	//1.接收传参,后端效验
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidParam)
		return
	}
	//2.业务逻辑
	m, err := logic.GetModuleDetail(id)
	if err != nil {
		zap.L().Error("logic.GetModuleDetail failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, m)
}

// UserGetModuleDetailHandler 用户获取模块详情
// @Summary 用户获取模块详情
// @Description 用户只可以获取公共模块的详情和属于自己的模块的详情
// @Tags 前台相关接口
// @Produce application/json
// @Param id path int true "模块id"
// @Param Authorization header string true "Bearer 用户token"
// @Success 200 {object} _RespModuleDetail
// @Router /api/v1/module/{id} [get]
func UserGetModuleDetailHandler(c *gin.Context) {
	//1.接收传参
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidParam)
		return
	}
	//2.后端效验：该模块是否是公共模块或属于该用户
	userid, exist := c.Get(CtxUserIDKey)
	if !exist {
		zap.L().Error("userid don't exist")
		RespErr(c, CodeServerBusy)
		return
	}
	if false == validate.IsCommonModuleOrMyModule(userid.(int64), id) {
		zap.L().Error("false == validate.CommonModuleOrMyModule(userid,id)")
		RespErr(c, CodeServerBusy)
		return
	}
	//3.业务逻辑
	m, err := logic.GetModuleDetail(id)
	if err != nil {
		zap.L().Error("logic.GetModuleDetail failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, m)
}

//UpdateModuleHandler 更新模块
func UpdateModuleHandler(c *gin.Context) {
	//1.接收参数，后端效验
	var params models.ParamsUpdateModule
	if err := ValidateJSONParam(c, &params); err != nil {
		zap.L().Error("update modules with invalid param", zap.Error(err))
		return
	}
	//2.业务逻辑
	if err := logic.UpdateModule(params); err != nil {
		zap.L().Error("logic.UpdateModule failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, nil)
}

//UpdateMyModuleHandler 更新我的模块
func UpdateMyModuleHandler(c *gin.Context) {
	//1.接收参数，后端效验
	var params models.ParamsUpdateModule
	if err := ValidateJSONParam(c, &params); err != nil {
		zap.L().Error("update modules with invalid param", zap.Error(err))
		return
	}
	//效验该模块是否属于该用户
	userID, exist := c.Get(CtxUserIDKey)
	if !exist {
		zap.L().Error("userid don't exist")
		RespErr(c, CodeServerBusy)
		return
	}
	if false == mysql.ValidateModuleOwner(userID.(int64), params.ID) {
		zap.L().Error("false == mysql.ValidateModuleOwner(userID, params.ID) ")
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑
	if err := logic.UpdateModule(params); err != nil {
		zap.L().Error("logic.UpdateModule failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, nil)
}

//DeleteModuleHandler 删除模块
func DeleteModuleHandler(c *gin.Context) {
	//1.接收参数，后端效验
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidParam)
		return
	}
	//2.业务逻辑
	err = logic.DeleteModule(id)
	if err != nil {
		zap.L().Error("logic.DeleteModule failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, nil)
}

//DeleteMyModuleHandler 删除我的模块
func DeleteMyModuleHandler(c *gin.Context) {
	//1.接收参数，后端效验
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidParam)
		return
	}
	userID, exist := c.Get(CtxUserIDKey)
	if !exist {
		zap.L().Error("userid don't exist")
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑
	err = logic.DeleteMyModule(id, userID.(int64))
	if err != nil {
		zap.L().Error("logic.DeleteMyModule failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, nil)
}

//GetModuleNoticeHandler 获取添加的模块的注意事项
func GetModuleNoticeHandler(c *gin.Context) {
	notice := fmt.Sprintf("服务器端默认情况下只接受xss_payload的get传参以及content-type为application/x-www-form-urlencoded的post传参。在编写xxs_payload时,用%s来代表本平台的url地址,用%s代表数据应该发送给哪个项目,用%s代表某某配置项的值",
		viper.GetString("payload.baseurl"), viper.GetString("payload.url_key"), viper.GetString("payload.option"))
	RespSuc(c, notice)
}

//GetMyModulesHandler 获取我的模块
func GetMyModulesHandler(c *gin.Context) {
	//1.接收参数,后端效验
	offset, err := strconv.ParseInt(c.Param("offset"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidOffset)
		return
	}
	count, err := strconv.ParseInt(c.Param("count"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidCount)
		return
	}
	userID, exist := c.Get(CtxUserIDKey)
	if !exist {
		zap.L().Error("c.Get(CtxUserIDKey) failed", zap.Error(err))
		return
	}
	username, exist := c.Get(CtxUsernameKey)
	if !exist {
		zap.L().Error("c.Get(CtxUsernameKey) failed", zap.Error(err))
		return
	}
	//2.业务逻辑
	modules, err := logic.GetMyModules(offset, count, userID.(int64), username.(string))
	if err != nil {
		zap.L().Error("logic.GetCommonModules failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	//3.返回响应
	RespSuc(c, modules)
}
