package filesystem

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/pkg/consts/content"
)

type Entity struct {
	Path        string
	Type        Type
	ContentType content.Type
}

func (f *Entity) Print(identation string) {
	fmt.Printf("%s%s, %s, %s\n", identation, f.Type, f.ContentType, f.Path)
}
