package performance

import (
	"database/sql"
	"encoding/json"
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

func (obj *Explain) Analyze(sql string) (en *ExplainEntity, s string, e error) {
	rows := obj.SelectAll("explain " + sql)
	defer rows.Close()
	var entity ExplainEntity
	for rows.Next() {
		e := rows.Scan(
			&entity.Id,
			&entity.Select_type,
			&entity.Table,
			&entity.Partitions,
			&entity.Type,
			&entity.Possible_keys,
			&entity.Key,
			&entity.Key_len,
			&entity.Ref,
			&entity.Rows,
			&entity.Filtered,
			&entity.Extra,
		)
		if e != nil {
			fmt.Println(e)
			return nil, "", e
		}
		jsonBytes, err := json.Marshal(entity)
		if err != nil {
			fmt.Println(err)
			return nil, "", err
		}
		return &entity, string(jsonBytes), nil
	}
	return nil, "", nil
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
