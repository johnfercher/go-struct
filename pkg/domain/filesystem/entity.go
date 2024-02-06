package filesystem

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/content"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/file"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/entities"
)

type Entity struct {
	Name        string
	Path        string
	Type        file.Type
	ContentType content.Type
	Structs     []*entities.Struct
	Interfaces  []*entities.Interface
}

func (f *Entity) Print(identation string) {
	fmt.Println("")
	//fmt.Printf("%s%s, %s, %s, %s\n", identation, f.Type, f.ContentType, f.Name, f.Path)
	if f.Type == file.Dir {
		fmt.Printf("%s/%s\n", identation, f.Name)
	} else {
		fmt.Printf("%s%s\n", identation, f.Name)
	}

	fmt.Println("STRUCTS")
	for _, str := range f.Structs {
		fmt.Println(str.String())
	}

	fmt.Println("INTERFACES")
	for _, inter := range f.Interfaces {
		fmt.Println(inter.String())
	}
}
