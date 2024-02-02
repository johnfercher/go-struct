package classifiers

import (
	"errors"
	"regexp"
	"strings"
)

var GoPackageRegex = regexp.MustCompile(`package \w+`)
var GoStructRegex = regexp.MustCompile(`type \w+ struct`)
var GoInterfaceRegex = regexp.MustCompile(`type \w+ interface`)
var GoFunctionRegex = regexp.MustCompile(`\w+()`)
var InterfaceRegex = regexp.MustCompile(`type\s+\w+\s+interface\s*{[^}]*}`)
var StructRegex = regexp.MustCompile(`type\s+\w+\s+struct\s*{[^}]*}`)

func GetGoInterface(content string) string {
	content = strings.ReplaceAll(content, "interface{}", "RAW_INTERFACE")
	name := InterfaceRegex.FindString(content)
	return name
}

func GetGoStruct(content string) string {
	return InterfaceRegex.FindString(content)
}

func GetGoPackageName(content string) (string, error) {
	name := GoPackageRegex.FindString(content)
	if name == "" {
		return "", errors.New("invalid package")
	}

	name = strings.ReplaceAll(name, "package ", "")
	return name, nil
}

func GetGoInterfacesName(content string) ([]string, error) {
	interfacesName := GoInterfaceRegex.FindAllString(content, -1)
	if len(interfacesName) == 0 {
		return nil, errors.New("invalid package")
	}

	for i := 0; i < len(interfacesName); i++ {
		interfacesName[i] = strings.ReplaceAll(interfacesName[i], "type ", "")
		interfacesName[i] = strings.ReplaceAll(interfacesName[i], " interface", "")
	}

	return interfacesName, nil
}

func GetGoStructsName(content string) ([]string, error) {
	structsName := GoStructRegex.FindAllString(content, -1)
	if len(structsName) == 0 {
		return nil, errors.New("invalid package")
	}

	for i := 0; i < len(structsName); i++ {
		structsName[i] = strings.ReplaceAll(structsName[i], "type ", "")
		structsName[i] = strings.ReplaceAll(structsName[i], " struct", "")
	}

	return structsName, nil
}
