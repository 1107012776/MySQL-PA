package mysql_pa

import (
	"database/sql"
	"fmt"
)

type Explain struct {
}

func (obj *Explain) Analyze() (b bool) {
	fmt.Println("analyze")
	return true
}

func (obj *Explain) GetDb() (db *sql.DB) {
	fmt.Println("getDb")
	return GetDb("phpshardingpdo1")
}
