package services

import (
	"github.com/johnfercher/go-pkg-struct/pkg/domain/entities"
	"github.com/johnfercher/go-pkg-struct/pkg/services/regex"
	"strings"
)

type StructInterpreter interface {
	ParseAll(content string) []*entities.Struct
}

type structInterpreter struct {
}

func NewStructInterpreter() *structInterpreter {
	return &structInterpreter{}
}

func (s *structInterpreter) ParseAll(content string) []*entities.Struct {
	packageName := regex.ExtractPackageName(content)
	imports := regex.ExtractImports(content)
	structContents := s.ExtractStructs(content)

	if len(structContents) == 0 {
		return nil
	}

	var structs []*entities.Struct
	for _, structContent := range structContents {
		structName := s.ExtractStructName(structContent)
		structs = append(structs, &entities.Struct{
			Package: packageName,
			Name:    structName,
			Imports: imports,
			Fields:  s.ExtractFields(structContent, imports),
			Methods: s.ExtractMethods(content, packageName, structName),
		})
	}

	return structs
}

func (s *structInterpreter) ExtractStructs(content string) []string {
	return regex.GoStruct.FindAllString(content, -1)
}

func (s *structInterpreter) ExtractStructName(content string) string {
	name := regex.GoStructName.FindString(content)
	name = strings.ReplaceAll(name, "type ", "")
	return strings.ReplaceAll(name, " struct", "")
}

func (s *structInterpreter) ExtractFields(content string, imports []*entities.Import) []*entities.Field {
	lines := strings.Split(content, "\n")
	lines = lines[1 : len(lines)-1]

	var fields []*entities.Field
	for _, line := range lines {
		fields = append(fields, &entities.Field{
			Content: strings.ReplaceAll(line, "\t", ""),
			Imports: s.getImportsMatched(line, imports),
		})
	}

	return fields
}

func (s *structInterpreter) ExtractMethods(content string, pkg string, structName string) []*entities.Function {
	methods := regex.GoStructMethods.FindAllString(content, -1)

	var functions []*entities.Function
	for i := 0; i < len(methods); i++ {
		methods[i] = regex.GoStructMethodsReceiver.ReplaceAllLiteralString(methods[i], "")

		methodName := regex.InterfaceMethodName.FindString(methods[i])
		methodName = strings.ReplaceAll(methodName, "(", "")
		argIn := regex.InArg.FindString(methods[i])
		argIn = strings.ReplaceAll(argIn, "(", "")
		argIn = strings.ReplaceAll(argIn, ") ", "")
		argOut := regex.OutArg.FindString(methods[i])
		argOut = strings.ReplaceAll(argOut, "( ", "")
		argOut = strings.ReplaceAll(argOut, ") ", "")

		functions = append(functions, &entities.Function{
			Package: pkg,
			Name:    methodName,
			Struct:  structName,
		})
	}

	return functions
}

func (s *structInterpreter) getImportsMatched(argString string, imports []*entities.Import) []*entities.Import {
	var usedImports []*entities.Import
	for _, imp := range imports {
		if imp.IsUsedIn(argString) {
			usedImports = append(usedImports, imp)
		}
	}
	return usedImports
}
