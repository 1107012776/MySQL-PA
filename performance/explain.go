package performance

import (
	"database/sql"
	"fmt"
)

type Explain struct {
	Db *sql.DB
}

func (obj *Explain) Analyze() (b bool) {
	fmt.Println("analyze")
	return true
}

func (obj *Explain) GetDb(name string, configPath string) (db *sql.DB) {
	obj.Db = GetDb(name, configPath)
	return obj.Db
}
