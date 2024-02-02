package entities

type Arg struct {
	Name string
	Type string
}

type Function struct {
	Package string
	Name    string
	Struct  string
	In      []Arg
	Out     []Arg
}

type Interface struct {
	Package string
	Name    string
	In      []Arg
	Out     []Arg
}

type Struct struct {
	Package    string
	Name       string
	Methods    []Function
	Interfaces []Interface
}
