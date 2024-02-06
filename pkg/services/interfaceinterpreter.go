package services

import (
	"github.com/johnfercher/go-pkg-struct/pkg/domain/entities"
	"github.com/johnfercher/go-pkg-struct/pkg/services/regex"
	"strings"
)

type InterfaceInterpreter interface {
	ParseAll(content string) []*entities.Interface
}

type interfaceInterpreter struct {
}

func NewInterfaceInterpreter() InterfaceInterpreter {
	return &interfaceInterpreter{}
}

func (i *interfaceInterpreter) ParseAll(content string) []*entities.Interface {
	packageName := regex.ExtractPackageName(content)
	imports := regex.ExtractImports(content)
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
			Methods: i.ExtractInterfaceMethods(interfaceContent, imports),
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

func (i *interfaceInterpreter) ExtractInterfaceMethods(content string, imports []*entities.Import) []*entities.Function {
	lines := strings.Split(content, "\n")
	methodsString := lines[1 : len(lines)-1]

	var methods []*entities.Function
	for index := 0; index < len(methodsString); index++ {
		methodString := strings.ReplaceAll(methodsString[index], "\t", "")
		methodName := regex.InterfaceMethodName.FindString(methodString)
		methodName = strings.ReplaceAll(methodName, "(", "")
		argIn := regex.InArg.FindString(methodString)
		argIn = strings.ReplaceAll(argIn, "(", "")
		argIn = strings.ReplaceAll(argIn, ") ", "")
		argOut := regex.OutArg.FindString(methodString)
		argOut = strings.ReplaceAll(argOut, "( ", "")
		argOut = strings.ReplaceAll(argOut, ") ", "")

		argsInString := strings.Split(argIn, ",")
		var argsIn []*entities.Field
		for _, argInString := range argsInString {
			argsIn = append(argsIn, &entities.Field{
				Content: argInString,
				Imports: i.getImportsMatched(argInString, imports),
			})
		}

		argsOutString := strings.Split(argOut, ",")
		var argsOut []*entities.Field
		for _, argOutString := range argsOutString {
			argsOut = append(argsOut, &entities.Field{
				Content: argOutString,
				Imports: i.getImportsMatched(argOutString, imports),
			})
		}

		methods = append(methods, &entities.Function{
			Name: methodName,
			In:   argsIn,
			Out:  argsOut,
		})
	}

	return methods
}

func (i *interfaceInterpreter) getImportsMatched(argString string, imports []*entities.Import) []*entities.Import {
	var usedImports []*entities.Import
	for _, imp := range imports {
		if imp.IsUsedIn(argString) {
			usedImports = append(usedImports, imp)
		}
	}
	return usedImports
}
