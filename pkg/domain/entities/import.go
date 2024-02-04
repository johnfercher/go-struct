package entities

import (
	"fmt"
	"strings"
)

type Import struct {
	Path    string
	Package string
}

func (i *Import) String() string {
	return fmt.Sprintf("%s: %s", i.Package, i.Path)
}

func (i *Import) IsUsedIn(line string) bool {
	return strings.Contains(line, i.Package)
}
