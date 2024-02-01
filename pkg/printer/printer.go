package printer

import (
	"github.com/johnfercher/go-pkg-struct/pkg/discover"
	"github.com/johnfercher/go-tree/node"
)

func Print(n *node.Node[discover.FileDir]) {
	identation := ""
	print(identation, n)
}

func print(identation string, n *node.Node[discover.FileDir]) {
	fileDir := n.GetData()
	fileDir.Print(identation)
	identation += " "

	nexts := n.GetNexts()

	for _, next := range nexts {
		print(identation, next)
	}
}
