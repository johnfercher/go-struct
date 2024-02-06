package entities

import "fmt"

type Function struct {
	Package string
	Name    string
	Struct  string
	In      []*Field
	Out     []*Field
}

func (f *Function) String() string {
	return fmt.Sprintf("%s(%s) %s", f.Name, f.getInString(), f.getOutString())
}

func (f *Function) getInString() string {
	var args string
	for _, arg := range f.In {
		args += arg.String() + " "
	}

	return args
}

func (f *Function) getOutString() string {
	var args string
	for _, arg := range f.Out {
		args += arg.String() + " "
	}

	return args
}
