package discover

import "fmt"

type Type string

const (
	File Type = "file"
	Dir  Type = "dir"
)

type FileDir struct {
	Path string
	Type Type
}

func (f *FileDir) Print(identation string) {
	fmt.Printf("%s%s: %s\n", identation, f.Type, f.Path)
}
