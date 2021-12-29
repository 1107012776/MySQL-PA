package mysql

import "fmt"

type Explain struct {
}

func (obj *Explain) Analyze() (b bool) {
	fmt.Println("analyze")
	return true
}
