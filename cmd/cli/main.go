package main

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/pkg/services"
	"github.com/johnfercher/go-pkg-struct/pkg/services/discover"
	"github.com/johnfercher/go-pkg-struct/pkg/services/loader"
	"log"
	"os"
)

func main() {
	packagePath := "docs/api"
	loader := loader.New()
	classifier := services.New()
	discover := discover.New(loader, classifier)
	interfaceInterpreter := services.NewInterfaceInterpreter()
	structInterpreter := services.NewStructInterpreter()

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := workdir + "/" + packagePath
	fmt.Println(path)

	entities, err := discover.Project(path)
	if err != nil {
		log.Fatal(err)
	}

	for key, entity := range entities {
		content, _ := loader.File(key)
		interfaces := interfaceInterpreter.ParseAll(string(content), key)
		structs := structInterpreter.ParseAll(string(content), key)
		newEntity := entity
		newEntity.Interfaces = interfaces
		newEntity.Structs = structs
		entities[key] = newEntity
	}

	for _, entity := range entities {
		entity.Print("")
	}
}
