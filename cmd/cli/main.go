package main

import (
	"fmt"
	"github.com/johnfercher/go-pkg-struct/pkg/classifier"
	"github.com/johnfercher/go-pkg-struct/pkg/discover"
	"github.com/johnfercher/go-pkg-struct/pkg/loader"
	"github.com/johnfercher/go-pkg-struct/pkg/printer"
	"log"
	"os"
)

func main() {
	packagePath := "docs/api"
	l := loader.New()
	c := classifier.New()
	d := discover.New(l, c)

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := workdir + "/" + packagePath
	fmt.Println(path)

	node, err := d.FindDirs(path)
	if err != nil {
		log.Fatal(err)
	}

	printer.Print(node)
}
