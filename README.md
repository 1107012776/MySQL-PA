# MySQL-Performance-Analysis
MySQL性能分析工具 (Mysql Performance Analysis)

# 示例（tests目录）
```golang
package tests

import (
	"fmt"
	"github.com/1107012776/MySQL-Performance-Analysis/performance"
	assert "github.com/magiconair/properties/assert"
	"testing"
)
func Test_AnalyzeIdSearch(t *testing.T) {
	p := performance.Explain{}
	p.Init("phpshardingpdo1", "./config.json")
	entity, err := p.Analyze("SELECT * from article_1 where id = 4;")
	assert.Equal(t, err == nil, true)
	assert.Equal(t, entity != nil, true)
	fmt.Println(entity.SelectType.String)
	fmt.Println(entity.ToJson())
}
```
