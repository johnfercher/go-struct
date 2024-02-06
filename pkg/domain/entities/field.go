package entities

import "fmt"

type Field struct {
	Content string
	Imports []*Import
}

func (a *Field) String() string {
	var imports string
	for _, imp := range a.Imports {
		imports += imp.String() + " "
	}
	return fmt.Sprintf("%s --%s--", a.Content, a.Imports)
}
