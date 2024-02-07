package entities

import "fmt"

type Function struct {
	Package      string
	Name         string
	Struct       string
	IsEntrypoint bool
	In           []*Field
	Out          []*Field
}

func (f *Function) String() string {
	return fmt.Sprintf("%s(%s) %s entrypoint(%v)", f.Name, f.getInString(), f.getOutString(), f.IsEntrypoint)
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
