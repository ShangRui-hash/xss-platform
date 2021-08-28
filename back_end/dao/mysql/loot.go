package mysql

import "xss/models"

//AddLoot 添加战利品
func AddLoot(Loot *models.Loot) error {
	s := `INSERT INTO loot(url_key,content) VALUES(?,?)`
	result, err := db.Exec(s, Loot.URLKey, Loot.Content)
	if err != nil {
		return err
	}
	Loot.ID, err = result.LastInsertId()
	return err
}

//GetLoots 获取项目所有战利品
func GetLoots(URLKey string) (loots []models.Loot, err error) {
	s := `SELECT * FROM loot WHERE url_key=?`
	err = db.Select(&loots, s, URLKey)
	return loots, err
}

//GetLoot 获取一个战利品
func GetLoot(lootID int64) (loot models.Loot, err error) {
	s := `SELECT * FROM loot WHERE id=?`
	err = db.Get(&loot, s, lootID)
	return loot, err
}

//GetLootCount 获取战利品数
func GetLootCount(URLKey string) (count int, err error) {
	s := `SELECT count(*) FROM loot WHERE url_key=?`
	err = db.Get(&count, s, URLKey)
	return count, err
}

//DelMyLoot 删除我的战利品
func DelMyLoot(id int64) error {
	_, err := db.Exec("DELETE FROM loot WHERE id = ?", id)
	return err
}

//ValidateLootOwner 效验loot的拥有者
func ValidateLootOwner(userID, lootID int64) bool {
	loot, err := GetLoot(lootID)
	if err != nil {
		return false
	}
	ownerID, err := GetOwnerOfProject(loot.URLKey)
	if err != nil {
		return false
	}
	return ownerID == userID
}
