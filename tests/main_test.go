package tests

import (
	"MySQL-PA/mysql_pa"
	"fmt"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_Analyze(t *testing.T) {
	p := mysql_pa.Explain{}
	assert.Equal(t, p.Analyze(), true)
}

func Test_GetDb(t *testing.T) {
	p := mysql_pa.Explain{}
	db := p.GetDb("phpshardingpdo1", "./config.json")
	fmt.Println(db)
	assert.Equal(t, db != nil, true)
}
