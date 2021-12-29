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
	assert.Equal(t, db.Ping() == nil, true)
}
