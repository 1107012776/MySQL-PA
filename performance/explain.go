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
	Id           sql.NullString `json:"id":`
	SelectType   sql.NullString `json:"select_type":`
	Table        sql.NullString `json:"table":`
	Partitions   sql.NullString `json:"partitions":`
	Type         sql.NullString `json:"type":`
	PossibleKeys sql.NullString `json:"possible_keys":`
	Key          sql.NullString `json:"key":`
	KeyLen       sql.NullString `json:"key_len":`
	Ref          sql.NullString `json:"ref":`
	Rows         sql.NullString `json:"rows":`
	Filtered     sql.NullString `json:"filtered":`
	Extra        sql.NullString `json:"Extra":`
}

func (entity *ExplainEntity) ToJson() (string, error) {
	jsonBytes, err := json.Marshal(entity)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(jsonBytes), nil
}

func (obj *Explain) Analyze(sql string) (*ExplainEntity, error) {
	rows := obj.SelectAll("explain " + sql)
	defer rows.Close()
	var entity ExplainEntity
	for rows.Next() {
		e := rows.Scan(
			&entity.Id,
			&entity.SelectType,
			&entity.Table,
			&entity.Partitions,
			&entity.Type,
			&entity.PossibleKeys,
			&entity.Key,
			&entity.KeyLen,
			&entity.Ref,
			&entity.Rows,
			&entity.Filtered,
			&entity.Extra,
		)
		if e != nil {
			fmt.Println(e)
			return nil, e
		}
		return &entity, nil
	}
	return nil, nil
}

func (obj *Explain) GetDb(name string, configPath string) *sql.DB {
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
