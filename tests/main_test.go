package tests

import (
	"MySQL-Performance-Analysis/performance"
	"encoding/json"
	"fmt"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_Analyze(t *testing.T) {
	p := performance.Explain{}
	assert.Equal(t, p.Analyze(), true)
}

func Test_GetDb(t *testing.T) {
	p := performance.Explain{}
	db := p.GetDb("phpshardingpdo1", "./config.json")
	fmt.Println(db)
	fmt.Println(p.Db)
	assert.Equal(t, db.Ping() == nil, true)
}

func Test_explain(t *testing.T) {
	p := performance.Explain{}
	p.GetDb("phpshardingpdo1", "./config.json")
	rows := p.SelectAll("show tables;")
	defer rows.Close()
	var table string
	for rows.Next() {
		e := rows.Scan(&table)
		columns, _ := rows.Columns()
		fmt.Println(len(columns))
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(table)
	}

}

func Test_select(t *testing.T) {
	p := performance.Explain{}
	p.GetDb("phpshardingpdo1", "./config.json")
	rows := p.SelectAll("explain SELECT * from article_1 where article_title = '张三是某网络科技的呀';")
	defer rows.Close()
	var entity performance.ExplainEntity
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
		columns, _ := rows.Columns()
		fmt.Println(columns)
		if e != nil {
			fmt.Println(e)
		}

		jsonBytes, err := json.Marshal(entity)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(jsonBytes))

		fmt.Println(entity)
	}

}
