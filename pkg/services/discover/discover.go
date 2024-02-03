package discover

import (
	"errors"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/content"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/file"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/filesystem"
	"github.com/johnfercher/go-pkg-struct/pkg/services"
	"github.com/johnfercher/go-pkg-struct/pkg/services/loader"
	"github.com/johnfercher/go-tree/node"
	"log"
	"os"
)

const goMod = "/go.mod"

type Discover interface {
	Project(path string) (*node.Node[filesystem.Entity], error)
}

type discover struct {
	loader         loader.Loader
	fileClassifier services.FileClassifier
}

func New(loader loader.Loader, fileClassifier services.FileClassifier) Discover {
	return &discover{
		loader:         loader,
		fileClassifier: fileClassifier,
	}
}

func (d *discover) Project(path string) (*node.Node[filesystem.Entity], error) {
	root := node.New(filesystem.Entity{
		Path: "root",
		Type: file.Dir,
	})

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := file.File
		if e.IsDir() {
			innerFileDirType = file.Dir
		}

		if innerFileDirType == file.Dir {
			inner, err := d.findDir(path+"/"+e.Name(), e.Name(), innerFileDirType)
			if err != nil {
				return nil, err
			}
			root.AddNext(inner)
		} else {
			filePath := path + "/" + e.Name()
			fileContent, err := d.loader.File(filePath)
			if err != nil {
				return nil, err
			}
			node := node.New(filesystem.Entity{
				Name:        e.Name(),
				Path:        filePath,
				Type:        file.File,
				ContentType: d.fileClassifier.Classify(string(fileContent)),
			})
			root.AddNext(node)
		}

	}

	node, ok := root.Filter(func(obj filesystem.Entity) bool {
		return obj.Type == file.Dir || obj.ContentType == content.Go
	})
	if !ok {
		log.Fatal(errors.New("there is no go files in any dir"))
	}

	return node, nil
}

func (d *discover) findDir(path string, name string, fileDirType file.Type) (*node.Node[filesystem.Entity], error) {
	root := node.New(filesystem.Entity{
		Path: path,
		Name: name,
		Type: fileDirType,
	})

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := file.File
		if e.IsDir() {
			innerFileDirType = file.Dir
		}

		if innerFileDirType == file.Dir {
			inner, err := d.findDir(path+"/"+e.Name(), e.Name(), innerFileDirType)
			if err != nil {
				return nil, err
			}
			root.AddNext(inner)
		} else {
			filePath := path + "/" + e.Name()
			fileContent, err := d.loader.File(filePath)
			if err != nil {
				return nil, err
			}
			node := node.New(filesystem.Entity{
				Name:        e.Name(),
				Path:        filePath,
				Type:        file.File,
				ContentType: d.fileClassifier.Classify(string(fileContent)),
			})
			root.AddNext(node)
		}

	}

	return root, nil
}
