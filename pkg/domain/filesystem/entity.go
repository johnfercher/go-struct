package filesystem

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/content"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/file"
)

type Entity struct {
	Name        string
	Path        string
	Type        file.Type
	ContentType content.Type
}

func (f *Entity) Print(identation string) {
	//fmt.Printf("%s%s, %s, %s, %s\n", identation, f.Type, f.ContentType, f.Name, f.Path)
	if f.Type == file.Dir {
		fmt.Printf("%s/%s\n", identation, f.Name)
	} else {
		fmt.Printf("%s%s\n", identation, f.Name)
	}

}
