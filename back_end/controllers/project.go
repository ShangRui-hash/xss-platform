package controllers

import (
	"strconv"
	"xss/logic"
	"xss/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetProjectFormHandler 获取创建项目/配置项目需要填的表单
// @Summary 获取创建项目/配置项目需要填的表单
// @Description 获取添加项目需要的表单，因为模块是变化的，所以添加项目需要的填的表单项需要后端动态生成。创建项目时需要填的是一个空表单，无序提交id。配置项目时需要对有数据的表单进行修改，需要提交项目id
// @Tags 前台相关接口
// @Accept application/json
// @Produce application/json
// @Param  Authorization header string true  "Bearer token"
// @Param object query int64  false "项目id"
// @Success 200 {object} _RespProjectForm
// @Router /api/v1/projectform [get]
func GetProjectFormHandler(c *gin.Context) {
	//接收参数
	userid, exist := c.Get(CtxUserIDKey)
	if !exist {
		RespErr(c, CodeServerBusy)
		return
	}
	projectID := c.Query("id")
	//业务逻辑
	forms, err := logic.GetCreateProjectForm(userid.(int64), projectID)
	if err != nil {
		zap.L().Error("logic.GetAllModules failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, forms)
}

//CreateProjectHandler 创建项目
func CreateProjectHandler(c *gin.Context) {
	//1.接收参数
	var params models.ParamCreateProject
	if err := ValidateJSONParam(c, &params); err != nil {
		zap.L().Error("create project with invalid param", zap.Error(err))
		return
	}
	userid, exist := c.Get(CtxUserIDKey)
	if !exist {
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑
	project, err := logic.CreateProject(params, userid.(int64))
	if err != nil {
		zap.L().Error("logic.CreateProject(params) failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, project)
}

//GetProjectListHandler 获取项目列表
func GetProjectListHandler(c *gin.Context) {
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
	userid, exist := c.Get(CtxUserIDKey)
	if !exist {
		RespErr(c, CodeServerBusy)
		return
	}
	//2.业务逻辑
	projects, err := logic.GetProjectList(offset, count, userid.(int64))
	if err != nil {
		zap.L().Error("logic.GetProjectList failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, projects)
}

//DeleteProjectHandler 删除项目
func DeleteProjectHandler(c *gin.Context) {
	//1.接收传参
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		RespErr(c, CodeInvalidParam)
		return
	}
	//2.业务逻辑
	if err := logic.DeleteProject(id); err != nil {
		zap.L().Error("logic.DeleteProject failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, nil)
}

//UpdateProjectHandler 更新项目
func UpdateProjectHandler(c *gin.Context) {
	//1.接收参数
	var params models.ParamCreateProject
	if err := ValidateJSONParam(c, &params); err != nil {
		zap.L().Error("create project with invalid param", zap.Error(err))
		return
	}

	userid, exist := c.Get(CtxUserIDKey)
	if !exist {
		RespErr(c, CodeServerBusy)
		return
	}
	URLKey := c.Param("url_key")
	if len(URLKey) != 4 {
		RespErr(c, CodeInvalidParam)
		return
	}
	//2.业务逻辑
	project, err := logic.UpdateProject(params, userid.(int64), URLKey)
	if err != nil {
		zap.L().Error("logic.UpdateProject(params) failed", zap.Error(err))
		RespErr(c, CodeServerBusy)
		return
	}
	RespSuc(c, project)
}
