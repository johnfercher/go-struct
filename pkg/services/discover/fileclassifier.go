package discover

import (
	"errors"
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/content"
	"github.com/johnfercher/go-pkg-struct/pkg/services/regex"
	"strings"
)

type FileClassifier interface {
	Classify(fileContent string) content.Type
}

type fileClassifier struct{}

func NewFileClassifier() *fileClassifier {
	return &fileClassifier{}
}

func (c *fileClassifier) Classify(fileContent string) content.Type {
	if c.IsGoFile(fileContent) {
		return content.Go
	}

	return content.Unknown
}

func (c *fileClassifier) IsGoFile(content string) bool {
	_, err := c.getGoPackageName(content)
	if err != nil {
		return false
	}

	return true
}

func (c *fileClassifier) getGoPackageName(content string) (string, error) {
	name := regex.GoPackageName.FindString(content)
	if name == "" {
		return "", errors.New("invalid package")
	}

	name = strings.ReplaceAll(name, "package ", "")
	return name, nil
}
