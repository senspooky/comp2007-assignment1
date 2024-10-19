package formatter

import "strings"

func (f formatter) Format(s string) string {
	var sb strings.Builder

	lines := lineBreakRegex.FindAllString(s, -1)
	if lines == nil {
		return ""
	}

	for _, line := range lines {
		leftPadding := 0
		switch f.alignment {
		case Centre:
			leftPadding = (maxWidth - len(line)) / 2
		case Right:
			leftPadding = (maxWidth - len(line))
		}
		sb.WriteString(strings.Repeat(" ", leftPadding))
		sb.WriteString(line)
		sb.WriteString("\n")
	}

	return sb.String()
}

func (f formatter) SetAlignment(a alignment) Formatter {
	if a.IsValid() {
		f.alignment = a
	}
	return f
}
