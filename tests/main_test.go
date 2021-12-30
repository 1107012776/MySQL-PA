package tests

import (
	"MySQL-Performance-Analysis/performance"
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
