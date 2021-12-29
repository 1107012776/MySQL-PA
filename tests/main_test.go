package tests

import (
	"MySQL-PA"
	assert "github.com/magiconair/properties/assert"
	"testing"
)

func Test_Analyze(t *testing.T) {
	p := MySQL_PA.Explain{}
	assert.Equal(t, p.Analyze(), true)
}

func Test_GetDb(t *testing.T) {
	p := MySQL_PA.Explain{}
	db := p.GetDb("phpshardingpdo1")
	assert.Equal(t, db != nil, true)
}
