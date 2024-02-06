package services

import (
	"fmt"
	"github.com/johnfercher/go-struct/pkg/domain/filesystem"
	"github.com/johnfercher/go-struct/pkg/services/discover"
	"github.com/johnfercher/go-struct/pkg/services/loader"
	"log"
	"os"
)

type ProjectInterpreter interface {
	Load(path string) (map[string]filesystem.Entity, error)
}

type projectInterpreter struct {
	loader               loader.Loader
	classifier           discover.FileClassifier
	discover             discover.Discover
	interfaceInterpreter InterfaceInterpreter
	structInterpreter    StructInterpreter
}

func New() ProjectInterpreter {
	loader := loader.New()
	classifier := discover.NewFileClassifier()
	discover := discover.New(loader, classifier)
	interfaceInterpreter := NewInterfaceInterpreter()
	structInterpreter := NewStructInterpreter()

	return &projectInterpreter{
		loader:               loader,
		classifier:           classifier,
		discover:             discover,
		interfaceInterpreter: interfaceInterpreter,
		structInterpreter:    structInterpreter,
	}
}

func (p *projectInterpreter) Load(packagePath string) (map[string]filesystem.Entity, error) {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := workdir + "/" + packagePath
	fmt.Println(path)

	entities, err := p.discover.Project(path)
	if err != nil {
		log.Fatal(err)
	}

	for key, entity := range entities {
		content, _ := p.loader.File(key)
		interfaces := p.interfaceInterpreter.ParseAll(string(content), key)
		structs := p.structInterpreter.ParseAll(string(content), key)
		newEntity := entity
		newEntity.Interfaces = interfaces
		newEntity.Structs = structs
		entities[key] = newEntity
	}

	for _, entity := range entities {
		entity.Print("")
	}

	return entities, nil
}
