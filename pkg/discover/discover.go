package discover

import (
	"github.com/johnfercher/go-pkg-struct/pkg/loader"
	"github.com/johnfercher/go-tree/node"
	"log"
	"os"
)

const goMod = "/go.mod"

type Discover interface {
	FindDirs(path string) *node.Node[FileDir]
}

type discover struct {
	loader loader.Loader
}

func New(loader loader.Loader) Discover {
	return &discover{
		loader: loader,
	}
}

func (d *discover) FindDirs(path string) *node.Node[FileDir] {
	root := node.New(FileDir{
		Path: "root",
		Type: Dir,
	})

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := File
		if e.IsDir() {
			innerFileDirType = Dir
		}

		if innerFileDirType == Dir {
			inner := d.findDir(path+"/"+e.Name(), innerFileDirType)
			root.AddNext(inner)
		} else {
			node := node.New(FileDir{
				Path: path + "/" + e.Name(),
				Type: File,
			})
			root.AddNext(node)
		}

	}

	return root
}

func (d *discover) findDir(path string, fileDirType Type) *node.Node[FileDir] {
	root := node.New(FileDir{
		Path: path,
		Type: fileDirType,
	})

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := File
		if e.IsDir() {
			innerFileDirType = Dir
		}

		if innerFileDirType == Dir {
			inner := d.findDir(path+"/"+e.Name(), innerFileDirType)
			root.AddNext(inner)
		} else {
			node := node.New(FileDir{
				Path: path + "/" + e.Name(),
				Type: File,
			})
			root.AddNext(node)
		}

	}

	return root
}
