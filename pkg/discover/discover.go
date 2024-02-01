package discover

import (
	"github.com/johnfercher/go-pkg-struct/pkg/classifier"
	"github.com/johnfercher/go-pkg-struct/pkg/consts/filesystem"
	"github.com/johnfercher/go-pkg-struct/pkg/loader"
	"github.com/johnfercher/go-tree/node"
	"log"
	"os"
)

const goMod = "/go.mod"

type Discover interface {
	FindDirs(path string) (*node.Node[filesystem.Entity], error)
}

type discover struct {
	loader     loader.Loader
	classifier classifier.Classifier
}

func New(loader loader.Loader, classifier classifier.Classifier) Discover {
	return &discover{
		loader:     loader,
		classifier: classifier,
	}
}

func (d *discover) FindDirs(path string) (*node.Node[filesystem.Entity], error) {
	root := node.New(filesystem.Entity{
		Path: "root",
		Type: filesystem.Dir,
	})

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := filesystem.File
		if e.IsDir() {
			innerFileDirType = filesystem.Dir
		}

		if innerFileDirType == filesystem.Dir {
			inner, err := d.findDir(path+"/"+e.Name(), innerFileDirType)
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
				Path:        filePath,
				Type:        filesystem.File,
				ContentType: d.classifier.Classify(string(fileContent)),
			})
			root.AddNext(node)
		}

	}

	return root, nil
}

func (d *discover) findDir(path string, fileDirType filesystem.Type) (*node.Node[filesystem.Entity], error) {
	root := node.New(filesystem.Entity{
		Path: path,
		Type: fileDirType,
	})

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		innerFileDirType := filesystem.File
		if e.IsDir() {
			innerFileDirType = filesystem.Dir
		}

		if innerFileDirType == filesystem.Dir {
			inner, err := d.findDir(path+"/"+e.Name(), innerFileDirType)
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
				Path:        filePath,
				Type:        filesystem.File,
				ContentType: d.classifier.Classify(string(fileContent)),
			})
			root.AddNext(node)
		}

	}

	return root, nil
}
