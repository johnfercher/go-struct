package main

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/pkg/services"
	"github.com/johnfercher/go-pkg-struct/pkg/services/discover"
	"github.com/johnfercher/go-pkg-struct/pkg/services/loader"
	"github.com/johnfercher/go-pkg-struct/pkg/services/printer"
	"log"
	"os"
)

func main() {
	packagePath := "docs/api"
	loader := loader.New()
	classifier := services.New()
	discover := discover.New(loader, classifier)

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := workdir + "/" + packagePath
	fmt.Println(path)

	node, err := discover.Project(path)
	if err != nil {
		log.Fatal(err)
	}

	printer.Print(node)
}
