package entities

import "fmt"

type Import struct {
	Path    string
	Package string
}

func (i *Import) String() string {
	return fmt.Sprintf("%s: %s", i.Package, i.Path)
}
