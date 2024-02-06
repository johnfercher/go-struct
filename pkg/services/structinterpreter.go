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
		structs = append(structs, &entities.Struct{
			Package: packageName,
			Name:    s.ExtractStructName(structContent),
			Imports: imports,
			Fields:  s.ExtractFields(content),
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

func (s *structInterpreter) ExtractFields(content string) []*entities.Field {
	return nil
}
