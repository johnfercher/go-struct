package discover

import (
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/file"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/filesystem"
	"github.com/johnfercher/go-pkg-struct/pkg/services"
	"github.com/johnfercher/go-pkg-struct/pkg/services/loader"
	"log"
	"os"
)

const goMod = "/go.mod"

type Discover interface {
	Project(path string) (map[string]filesystem.Entity, error)
}

type discover struct {
	loader         loader.Loader
	fileClassifier services.FileClassifier
	entities       map[string]filesystem.Entity
}

func New(loader loader.Loader, fileClassifier services.FileClassifier) Discover {
	return &discover{
		loader:         loader,
		fileClassifier: fileClassifier,
		entities:       make(map[string]filesystem.Entity),
	}
}

func (d *discover) Project(path string) (map[string]filesystem.Entity, error) {
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
			err := d.findDir(path+"/"+e.Name(), e.Name(), innerFileDirType)
			if err != nil {
				return nil, err
			}
		} else {
			filePath := path + "/" + e.Name()
			fileContent, err := d.loader.File(filePath)
			if err != nil {
				return nil, err
			}

			d.entities[path] = filesystem.Entity{
				Name:        e.Name(),
				Path:        filePath,
				Type:        file.File,
				ContentType: d.fileClassifier.Classify(string(fileContent)),
			}
		}
	}

	return d.entities, nil
}

func (d *discover) findDir(path string, name string, fileDirType file.Type) error {
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
			err := d.findDir(path+"/"+e.Name(), e.Name(), innerFileDirType)
			if err != nil {
				return err
			}
		} else {
			filePath := path + "/" + e.Name()
			fileContent, err := d.loader.File(filePath)
			if err != nil {
				return err
			}
			d.entities[filePath] = filesystem.Entity{
				Name:        e.Name(),
				Path:        filePath,
				Type:        file.File,
				ContentType: d.fileClassifier.Classify(string(fileContent)),
			}
		}

	}

	return nil
}
