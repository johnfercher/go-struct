package entities

type Function struct {
	Package string
	Name    string
	Struct  string
	In      []Arg
	Out     []Arg
}
