package entities

import (
	"fmt"
)

type Interface struct {
	Package string
	Name    string
	Path    string
	Imports []*Import
	Methods []*Function
}

func (i *Interface) String() string {
	var content string

	content += fmt.Sprintf("PACKAGE: %s\n", i.Package)
	content += fmt.Sprintf("NAME: %s\n", i.Name)
	content += fmt.Sprintf("PATH: %s\n", i.Path)

	content += fmt.Sprintf("IMPORTS: %s\n", i.Path)
	for _, imp := range i.Imports {
		content += imp.String() + "\n"
	}

	content += fmt.Sprintf("METHODS: %s\n", i.Path)
	for _, method := range i.Methods {
		content += method.String() + "\n"
	}

	return content
}
