package main

import (
	"MySQL-PA/mysql"
)

func main() {
	p := mysql.Explain{}
	p.Analyze()
}
