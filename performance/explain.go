package performance

import (
	"database/sql"
	"fmt"
)

type Explain struct {
	Db *sql.DB
}

type ExplainEntity struct {
	Id            string         `json:"id":`
	Select_type   string         `json:"select_type":`
	Table         string         `json:"table":`
	Partitions    sql.NullString `json:"partitions":`
	Type          string         `json:"type":`
	Possible_keys sql.NullString `json:"possible_keys":`
	Key           sql.NullString `json:"key":`
	Key_len       sql.NullInt64  `json:"key_len":`
	Ref           sql.NullString `json:"ref":`
	Rows          string         `json:"rows":`
	Filtered      string         `json:"filtered":`
	Extra         string         `json:"Extra":`
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
