package regex

import (
	"github.com/johnfercher/go-struct/pkg/domain/entities"
	"regexp"
	"strings"
)

var GoPackageName = regexp.MustCompile(`package \w+`)
var ImportWord = regexp.MustCompile(`import\s?\(?`)

var GoStructName = regexp.MustCompile(`type \w+ struct`)
var GoStruct = regexp.MustCompile(`type\s+\w+\s+struct\s*{[^}]*}`)
var GoStructFields = regexp.MustCompile(`\s+?\w+\s+(\[\])?\*?\w+(\.\w+)?`)
var GoStructMethods = regexp.MustCompile(`func \(\w+\s\*?\w+\)\s\w+\(.+\)(\s\w+(\.\w+)?)?`)
var GoStructMethodsReceiver = regexp.MustCompile(`func \(\w+\s\*?\w+\)\s`)

var GoInterfaceName = regexp.MustCompile(`type \w+ interface`)
var GoInterface = regexp.MustCompile(`type\s+\w+\s+interface\s*{[^}]*}`)
var InterfaceMethodName = regexp.MustCompile(`\w+\(`)

var InArg = regexp.MustCompile(`\(.+\) `)
var OutArg = regexp.MustCompile(`\) (\(.+\)|\w+)`)

func ExtractImports(content string) []*entities.Import {
	lines := strings.Split(content, "\n")

	var imports []*entities.Import
	for index := 0; index < len(lines); index++ {
		match := ImportWord.FindString(lines[index])
		if match == "import " {
			importString := strings.ReplaceAll(lines[index], "import ", "")
			return []*entities.Import{
				{
					Path:    importString,
					Package: getLastWord(importString),
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
					Package: getLastWord(line),
				})
			}
		}
	}

	return imports
}

func ExtractPackageName(content string) string {
	name := GoPackageName.FindString(content)
	return strings.ReplaceAll(name, "package ", "")
}

func getLastWord(line string) string {
	line = strings.ReplaceAll(line, "\"", "")
	words := strings.Split(line, "/")
	return words[len(words)-1]
}
