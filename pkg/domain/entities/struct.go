package entities

import "fmt"

type Struct struct {
	Package    string
	Name       string
	Path       string
	Imports    []*Import
	Methods    []*Function
	Interfaces []*Interface
	Fields     []*Field
}

func (s *Struct) String() string {
	var content string

	content += fmt.Sprintf("PACKAGE: %s\n", s.Package)
	content += fmt.Sprintf("NAME: %s\n", s.Name)
	content += fmt.Sprintf("PATH: %s\n", s.Path)

	content += fmt.Sprintf("IMPORTS: %s\n", s.Path)
	for _, imp := range s.Imports {
		content += imp.String() + "\n"
	}

	/*content += fmt.Sprintf("METHODS: %s\n", i.Path)
	for _, method := range s.Methods {
		content += method.String() + "\n"
	}*/

	return content
}
