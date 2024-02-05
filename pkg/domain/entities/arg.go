package entities

import "fmt"

type Arg struct {
	Content string
	Imports []*Import
}

func (a *Arg) String() string {
	var imports string
	for _, imp := range a.Imports {
		imports += imp.String() + " "
	}
	return fmt.Sprintf("%s --%s--", a.Content, a.Imports)
}
