package services

import (
	"github.com/johnfercher/go-pkg-struct/pkg/domain/entities"
	"github.com/johnfercher/go-pkg-struct/pkg/services/regex"
	"strings"
)

type InterfaceInterpreter interface {
	ParseAll(file string) []*entities.Interface
}

type interfaceInterpreter struct {
}

func NewInterfaceInterpreter() InterfaceInterpreter {
	return &interfaceInterpreter{}
}

func (i *interfaceInterpreter) ParseAll(content string) []*entities.Interface {
	packageName := i.ExtractPackageName(content)
	imports := i.ExtractImports(content)
	interfaceContents := i.ExtractInterfaces(content)
	if len(interfaceContents) == 0 {
		return nil
	}

	var interfaces []*entities.Interface
	for _, interfaceContent := range interfaceContents {
		interfaces = append(interfaces, &entities.Interface{
			Package: packageName,
			Name:    i.ExtractInterfaceName(interfaceContent),
			Imports: imports,
			Methods: i.ExtractInterfaceMethods(interfaceContent),
		})
	}

	return interfaces
}

func (i *interfaceInterpreter) ExtractInterfaces(content string) []string {
	content = strings.ReplaceAll(content, "interface{}", "RAW_INTERFACE") // Regex has no support to interface{} args in methods
	return regex.GoInterface.FindAllString(content, -1)
}

func (i *interfaceInterpreter) ExtractInterfaceName(content string) string {
	name := regex.GoInterfaceName.FindString(content)
	name = strings.ReplaceAll(name, "type ", "")
	return strings.ReplaceAll(name, " interface", "")
}

func (i *interfaceInterpreter) ExtractPackageName(content string) string {
	name := regex.GoPackageName.FindString(content)
	return strings.ReplaceAll(name, "package ", "")
}

func (i *interfaceInterpreter) ExtractInterfaceMethods(content string) []string {
	lines := strings.Split(content, "\n")
	methods := lines[1 : len(lines)-1]

	for i := 0; i < len(methods); i++ {
		methods[i] = strings.ReplaceAll(methods[i], "\t", "")
	}

	return methods
}

func (i *interfaceInterpreter) ExtractImports(content string) []*entities.Import {
	lines := strings.Split(content, "\n")

	var imports []*entities.Import
	for index := 0; index < len(lines); index++ {
		match := regex.ImportWord.FindString(lines[index])
		if match == "import " {
			importString := strings.ReplaceAll(lines[index], "import ", "")
			return []*entities.Import{
				&entities.Import{
					Path:    importString,
					Package: i.getLastWord(importString),
				},
			}
		}
		if match == "import (" {
			index++
			for j := index; j < len(lines); j++ {
				if lines[j] == ")" {
					break
				}
				line := strings.ReplaceAll(lines[j], "\t", "")
				imports = append(imports, &entities.Import{
					Path:    line,
					Package: i.getLastWord(line),
				})
			}
		}
	}

	return imports
}

func (i *interfaceInterpreter) getLastWord(line string) string {
	line = strings.ReplaceAll(line, "\"", "")
	words := strings.Split(line, "/")
	return words[len(words)-1]
}
