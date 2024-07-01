package pterm

var EscapePressed = tEscapePressed{}

type tEscapePressed struct {}

func (tEscapePressed) Error() string {
	return "escape pressed"
}
