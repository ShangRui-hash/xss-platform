package mysql

import "xss/models"

//CreateProject 创建项目
func CreateProject(project *models.Project) error {
	s := `INSERT INTO 
	project(name,description,user_id,url_key)
	VALUES (?,?,?,?)`
	r, err := db.Exec(s, project.Name, project.Desc, project.UserID, project.URLKey)
	if err != nil {
		return err
	}
	project.ID, err = r.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

//UpdateProject 更新项目
func UpdateProject(project *models.Project) error {
	s := `UPDATE project SET name=?,description=? WHERE user_id=? and url_key=?`
	_, err := db.Exec(s, project.Name, project.Desc, project.UserID, project.URLKey)
	return err
}

//GetProjectList 获取项目列表
func GetProjectList(offset, count, userid int64) (p []models.Project, err error) {
	s := `SELECT * FROM project WHERE user_id = ? LIMIT ?,?`
	err = db.Select(&p, s, userid, offset, count)
	return p, err
}

//GetProject 获取项目详情
func GetProject(id int64) (project models.Project, err error) {
	s := `SELECT * FROM project WHERE id = ?`
	err = db.Get(&project, s, id)
	return project, err
}

//GetOwnerOfProject 获取项目的主人
func GetOwnerOfProject(URLKey string) (userID int64, err error) {
	s := `SELECT user_id FROM project WHERE url_key=?`
	err = db.Get(&userID, s, URLKey)
	return userID, err
}

//DeleteProject 删除项目
func DeleteProject(id int64) error {
	s := `DELETE FROM project WHERE id=?`
	_, err := db.Exec(s, id)
	return err
}

//IsURLKeyExist 查询URLKey是否存在
func IsURLKeyExist(URLKey string) (bool, error) {
	sql := `SELECT count(id) FROM project WHERE url_key=?`
	var count int
	err := db.Get(&count, sql, URLKey)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
