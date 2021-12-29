package mysqlPa

import "fmt"

type Explain struct {
}

func (obj *Explain) Analyze() {
	fmt.Println("analyze")
}
