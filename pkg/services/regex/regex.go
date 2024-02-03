package regex

import (
	"regexp"
)

var GoPackageName = regexp.MustCompile(`package \w+`)
var GoStructName = regexp.MustCompile(`type \w+ struct`)
var GoInterfaceName = regexp.MustCompile(`type \w+ interface`)
var GoInterface = regexp.MustCompile(`type\s+\w+\s+interface\s*{[^}]*}`)
var StructRegex = regexp.MustCompile(`type\s+\w+\s+struct\s*{[^}]*}`)
var ImportWord = regexp.MustCompile(`import\s?\(?`)
