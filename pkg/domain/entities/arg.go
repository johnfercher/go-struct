package entities

import "fmt"

type Arg struct {
	Content string
	Import  []Import
}

func (a *Arg) String() string {
	return fmt.Sprintf("%s", a.Content)
}
