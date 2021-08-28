package logic

import (
	"fmt"
	"time"
	"xss/dao/mysql"
	"xss/dao/redis"
	"xss/models"
	"xss/pkg/randstr"

	"go.uber.org/zap"
)

//CreateProject 创建项目
func CreateProject(params models.ParamCreateProject, userid int64) (models.Project, error) {
	project := &models.Project{
		Name:      params.Name,
		Desc:      params.Desc,
		UserID:    userid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	//0.生成一个随机的URLKey
	URLKey, err := GenURLKey()
	if err != nil {
		return *project, err
	}
	project.URLKey = URLKey
	//1.项目入库
	if err := mysql.CreateProject(project); err != nil {
		return *project, err
	}
	//2.关联模块
	for i := range params.ModuleList {
		if params.ModuleList[i].IsChoosed == false {
			continue
		}
		err := AddModuleToProject(project.URLKey, params.ModuleList[i].ID, params.ModuleList[i].OptionList)
		if err != nil {
			return *project, err
		}
	}
	return *project, nil
}

//UpdateProject 更新项目
func UpdateProject(params models.ParamCreateProject, userID int64, URLKey string) (models.Project, error) {
	p := &models.Project{
		Name:   params.Name,
		URLKey: URLKey,
		Desc:   params.Desc,
		UserID: userID,
	}
	//更新基本信息
	if err := mysql.UpdateProject(p); err != nil {
		return *p, err
	}
	moduleSetOfProject := redis.NewModuleSetOfProject(p.URLKey)
	//更新项目和模块之间的关联
	for i := range params.ModuleList {
		exist, err := moduleSetOfProject.IsMember(fmt.Sprintf("%d", params.ModuleList[i].ID))
		if err != nil {
			return *p, err
		}

		if exist {
			if false == params.ModuleList[i].IsChoosed {
				//删除关联模块
				err := RemModuleFromProject(p.URLKey, params.ModuleList[i].ID)
				if err != nil {
					return *p, err
				}
			} else {
				//更新关联模块
				err := UpdateModuleOfProject(p.URLKey, params.ModuleList[i].ID, params.ModuleList[i].OptionList)
				if err != nil {
					return *p, err
				}
			}
		} else {
			//添加关联模块
			if true == params.ModuleList[i].IsChoosed {
				err := AddModuleToProject(p.URLKey, params.ModuleList[i].ID, params.ModuleList[i].OptionList)
				if err != nil {
					return *p, err
				}
			}
		}

	}
	return *p, nil
}

//AddModuleToProject 关联项目和模块
func AddModuleToProject(URLKey string, moduleID int64, OptionList []models.OptionItem) error {
	//添加关联
	_, err := redis.NewModuleSetOfProject(URLKey).Add(fmt.Sprintf("%d", moduleID))
	if err != nil {
		zap.L().Error(`moduleSetOfProject.Add failed`, zap.Error(err))
		return err
	}
	//关联项目和模块参数
	for _, option := range OptionList {
		_, err := redis.NewOptionOfProject(URLKey, moduleID).Set(option.Name, option.Value)
		if err != nil {
			zap.L().Error("reids.NewOptionOfProject.Set failed", zap.Error(err))
			return err
		}
	}
	return nil
}

//RemModuleFromProject 从项目中删除模块
func RemModuleFromProject(URLKey string, moduleID int64) error {
	//删除项目和模块之间的关联
	if _, err := redis.NewModuleSetOfProject(URLKey).Rem(fmt.Sprintf("%d", moduleID)); err != nil {
		zap.L().Error(`moduleSetOfProject.Rem failed`, zap.Error(err))
		return err
	}
	//删除项目和模块参数之间的关联
	if err := redis.NewOptionOfProject(URLKey, moduleID).Clear(); err != nil {
		zap.L().Error("redis.NewOptionOfProject.Clear failed", zap.Error(err))
		return err
	}
	return nil
}

//UpdateModuleOfProject 更新和项目关联的模块
func UpdateModuleOfProject(URLKey string, moduleID int64, OptionList []models.OptionItem) error {
	for _, option := range OptionList {
		_, err := redis.NewOptionOfProject(URLKey, moduleID).Set(option.Name, option.Value)
		if err != nil {
			zap.L().Error("reids.NewOptionOfProject.Set failed", zap.Error(err))
			return err
		}
	}
	return nil
}

//GetProjectList 获取项目列表
func GetProjectList(offset, count, userid int64) ([]models.Project, error) {
	//查询基本信息
	projects, err := mysql.GetProjectList(offset, count, userid)
	//查询战利品数量
	for i := range projects {
		projects[i].LootCount, err = mysql.GetLootCount(projects[i].URLKey)
		if err != nil {
			zap.L().Error("mysql.GetLootCount failed", zap.Error(err))
			return projects, err
		}
	}
	return projects, err

}

//DeleteProject 删除项目
func DeleteProject(id int64) error {
	//0.获取项目详情
	project, err := mysql.GetProject(id)
	if err != nil {
		zap.L().Error("mysql.GetProject failed")
		return err
	}
	//1.删除项目
	err = mysql.DeleteProject(id)
	if err != nil {
		zap.L().Error("mysql.DeleteProject failed")
		return err
	}
	//2.删除项目和模块之间的关联
	err = redis.NewModuleSetOfProject(project.URLKey).Clear()
	if err != nil {
		zap.L().Error("redis.NewModuleSetOfProject.Clear() failed", zap.Error(err))
		return err
	}
	//3.删除项目和模块参数之间的关联
	err = redis.ClearAllOptions(project.URLKey)
	if err != nil {
		zap.L().Error("redis.ClearAllOptions failed", zap.Error(err))
		return err
	}
	return nil
}

//GenURLKey 生成项目的URL_Key
func GenURLKey() (URLKey string, err error) {
	for {
		URLKey := randstr.String(4)
		exist, err := mysql.IsURLKeyExist(URLKey)
		if err != nil {
			return "", err
		}
		if exist == false {
			return URLKey, nil
		}
	}
}
