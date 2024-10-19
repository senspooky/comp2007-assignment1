package formatter

type alignment uint

const (
	Left alignment = iota
	Centre
	Right
)

func (a alignment) IsValid() bool {
	switch a {
	case Left:
		return true
	case Right:
		return true
	case Centre:
		return true
	}
	return false
}
