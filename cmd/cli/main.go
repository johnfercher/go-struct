package main

import (
	"github.com/johnfercher/go-pkg-struct/pkg/services"
	"log"
)

func main() {
	packagePath := "docs/api"
	projectInterpreter := services.New()
	entities, err := projectInterpreter.Load(packagePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range entities {
		entity.Print("")
	}
}
