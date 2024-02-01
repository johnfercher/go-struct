package classifier

import (
	"github.com/johnfercher/go-pkg-struct/pkg/consts/content"
	"regexp"
	"strings"
)

type Classifier interface {
	Classify(fileContent string) content.Type
}

type classifier struct {
	goRegex *regexp.Regexp
}

func New() *classifier {
	return &classifier{
		goRegex: regexp.MustCompile(`package \w+`),
	}
}

func (c *classifier) Classify(fileContent string) content.Type {
	if c.IsGo(fileContent) {
		return content.Go
	}

	return content.Unknown
}

func (c *classifier) IsGo(content string) bool {
	lines := strings.Split(content, "\n")
	if len(lines) == 0 {
		return false
	}

	return c.goRegex.MatchString(lines[0])
}
