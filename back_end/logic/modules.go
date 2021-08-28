package logic

import (
	"fmt"
	"strconv"
	"xss/dao/mysql"
	"xss/dao/redis"
	"xss/models"

	"github.com/syyongx/php2go"
	"go.uber.org/zap"
)

//GetModules 获取所有模块
func GetModules(offset, count int64) (moduleList []models.Module, err error) {
	return mysql.GetModules(offset, count)
}

//GetCommonModules 获取公共模块
func GetCommonModules(offset, count int64) (moduleList []models.Module, err error) {
	return mysql.GetCommonModules(offset, count)
}

//GetMyModules 获取我的模块
func GetMyModules(offset, count, userID int64, username string) (moduleList []models.Module, err error) {
	modules, err := mysql.GetMyModules(offset, count, userID)
	if err != nil {
		return modules, err
	}
	for i := range modules {
		modules[i].Username = username
	}
	return modules, nil
}

//AddModule 添加模块
func AddModule(params models.ParamsAddModule, userid int64, isadmin bool) (models.Module, error) {
	//1.名称、描述、代码等存入mysql
	m := models.Module{
		Name:       params.Name,
		Desc:       params.Desc,
		XSSPayload: params.XSSPayload,
		UserID:     userid,
		IsAdmin:    &isadmin,
		IsCommon:   params.IsCommon,
	}
	err := mysql.AddModule(&m)
	if err != nil {
		zap.L().Error("mysql.AddModule failed", zap.Error(err))
		return m, err
	}
	//2.参数列表、配置项列表存入redis
	if err := redis.NewParamSet(m.ID).Update(params.ParamList); err != nil {
		zap.L().Error("redis.NewParamSet.Update failed", zap.Error(err))
		return m, err
	}
	if err := redis.NewOptionSet(m.ID).Update(params.OptionList); err != nil {
		zap.L().Error("redis.NewOptionSet.Update failed", zap.Error(err))
		return m, err
	}
	return m, nil
}

//GetCreateProjectForm 获取创建项目的表单格式
func GetCreateProjectForm(userID int64, projectID string) (params models.ParamCreateProject, err error) {
	//获取所有模块
	modules, err := mysql.GetAllModules(userID)
	if err != nil {
		return params, err
	}
	params.ModuleList = make([]models.ModuleItem, 0, len(modules))
	//查询所有配置项
	for i := range modules {
		options, err := redis.NewOptionSet(modules[i].ID).Members()
		if err != nil {
			return params, err
		}
		optionList := make([]models.OptionItem, 0, len(options))
		for _, option := range options {
			optionList = append(optionList, models.OptionItem{
				Name:  option,
				Value: "",
			})
		}
		params.ModuleList = append(params.ModuleList, models.ModuleItem{
			ID:         modules[i].ID,
			Name:       modules[i].Name,
			IsChoosed:  false,
			OptionList: optionList,
		})
	}
	//如果没有指定project id 直接返回
	if len(projectID) == 0 {
		return params, nil
	}
	//查询project详情
	id, err := strconv.ParseInt(projectID, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		return params, err
	}
	project, err := mysql.GetProject(id)
	if err != nil {
		zap.L().Error("mysql.GetProject failed", zap.Error(err))
		return params, err
	}
	//如果越权直接返回
	if project.UserID != userID {
		return params, nil
	}
	params.Name = project.Name
	params.Desc = project.Desc
	//查询所有模块的选中情况
	moduleIDs, err := redis.NewModuleSetOfProject(project.URLKey).Members()
	if err != nil {
		zap.L().Error("redis.NewModuleSetOfProject failed", zap.Error(err))
		return params, err
	}
	//查询选中模块的参数情况
	for i, module := range params.ModuleList {
		if php2go.InArray(fmt.Sprintf("%d", module.ID), moduleIDs) == false {
			continue
		}
		params.ModuleList[i].IsChoosed = true
		options, err := redis.NewOptionOfProject(project.URLKey, module.ID).GetAll()
		if err != nil {
			zap.L().Error("redis.NewOptionOfProject failed", zap.Error(err))
			return params, err
		}
		for o := range params.ModuleList[i].OptionList {
			params.ModuleList[i].OptionList[o].Value = options[params.ModuleList[i].OptionList[o].Name]
		}
	}
	return params, nil
}

//GetModuleDetail 获取模块详情
func GetModuleDetail(id int64) (models.Module, error) {
	//获取基本信息
	m, err := mysql.GetModule(id)
	if err != nil {
		return m, err
	}
	//获取配置项
	m.OptionList, err = redis.NewOptionSet(id).Members()
	if err != nil {
		return m, err
	}
	//获取参数项
	m.ParamList, err = redis.NewParamSet(id).Members()
	if err != nil {
		return m, err
	}
	return m, nil
}

//UpdateModule 更新模块
func UpdateModule(params models.ParamsUpdateModule) error {
	if err := mysql.UpdateModule(params); err != nil {
		return err
	}
	err := redis.NewOptionSet(params.ID).Update(params.OptionList)
	if err != nil {
		return err
	}
	err = redis.NewParamSet(params.ID).Update(params.ParamList)
	if err != nil {
		return err
	}
	return nil
}

//DeleteModule 删除模块
func DeleteModule(id int64) error {
	if err := mysql.DeleteModule(id); err != nil {
		return err
	}
	err := redis.NewOptionSet(id).Clear()
	if err != nil {
		return err
	}
	err = redis.NewParamSet(id).Clear()
	if err != nil {
		return err
	}
	return nil
}

//DeleteMyModule 删除我的模块
func DeleteMyModule(moduleID, userID int64) error {
	if err := mysql.DeleteMyModule(moduleID, userID); err != nil {
		return err
	}
	err := redis.NewOptionSet(moduleID).Clear()
	if err != nil {
		return err
	}
	err = redis.NewParamSet(moduleID).Clear()
	if err != nil {
		return err
	}
	return nil
}
