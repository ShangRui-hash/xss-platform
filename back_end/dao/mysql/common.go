package mysql

import (
	"database/sql"
	"errors"
)

//ErrWrongAffectedRow 影响行数不为1
var ErrWrongAffectedRow = errors.New("rows affected != 1")

//ErrAdminNotExist 管理员数量不为1
var ErrAdminNotExist = errors.New("the count of admin name != 1")

//CheckRowsAffected 检查删除/修改操作是否成功并记录日志
func CheckRowsAffected(result sql.Result) error {
	n, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		return err
	}
	if n != 1 {
		return ErrWrongAffectedRow
	}
	return nil
}
