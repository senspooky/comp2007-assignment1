package formatter

import (
	"fmt"
	"regexp"
)

type Formatter interface {
	Format(string) string
	SetAlignment(alignment) Formatter
}

type formatter struct {
	alignment alignment
}

func F() Formatter {
	return &formatter{
		alignment: Left,
	}
}

var maxWidth int = 50

var lineBreakRegex = regexp.MustCompile(fmt.Sprintf(`.{1,%d}(?:\s|$)`, maxWidth))
