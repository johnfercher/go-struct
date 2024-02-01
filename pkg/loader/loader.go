package loader

import "os"

type Loader interface {
	File(file string) ([]byte, error)
}

type loader struct {
}

func New() Loader {
	return &loader{}
}

func (l *loader) File(file string) ([]byte, error) {
	return os.ReadFile(file)
}
