package performance

import (
	"database/sql"
	"fmt"
)

type Explain struct {
	Db *sql.DB
}

func (obj *Explain) Analyze() (b bool) {
	fmt.Println("analyze")
	return true
}

func (obj *Explain) GetDb(name string, configPath string) (db *sql.DB) {
	obj.Db = GetDb(name, configPath)
	return obj.Db
}

func (obj *Explain) SelectAll(sql string) *sql.Rows {
	db := obj.Db
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("查询出错了 %s", err)
		return nil
	}
	return rows
}
