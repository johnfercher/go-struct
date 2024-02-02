package classifiers

import (
	"github.com/johnfercher/go-pkg-struct/pkg/domain/consts/content"
)

type FileClassifier interface {
	Classify(fileContent string) content.Type
}

type fileClassifier struct{}

func New() *fileClassifier {
	return &fileClassifier{}
}

func (c *fileClassifier) Classify(fileContent string) content.Type {
	if c.IsGoFile(fileContent) {
		return content.Go
	}

	return content.Unknown
}

func (c *fileClassifier) IsGoFile(content string) bool {
	_, err := GetGoPackageName(content)
	if err != nil {
		return false
	}

	return true
}
