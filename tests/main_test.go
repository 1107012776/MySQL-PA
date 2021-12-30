package tests

import (
	"MySQL-Performance-Analysis/performance"
	"fmt"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_Analyze(t *testing.T) {
	p := performance.Explain{}
	p.Init("phpshardingpdo1", "./config.json")
	_, err := p.Analyze("SELECT * from article_1 where article_title = '张三是某网络科技的呀';")
	assert.Equal(t, err == nil, true)
}

func Test_AnalyzeIdSearch(t *testing.T) {
	p := performance.Explain{}
	p.Init("phpshardingpdo1", "./config.json")
	entity, err := p.Analyze("SELECT * from article_1 where id = 4;")
	assert.Equal(t, err == nil, true)
	assert.Equal(t, entity != nil, true)
	fmt.Println(entity.SelectType.String)
	fmt.Println(entity.ToJson())
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
	p.Init("phpshardingpdo1", "./config.json")
	rows, err := p.SelectAll("show tables;")
	assert.Equal(t, err == nil, true)
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
