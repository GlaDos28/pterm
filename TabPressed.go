package pterm

var TabPressed = tTabPressed{}

type tTabPressed struct {}

func (tTabPressed) Error() string {
	return "tab pressed"
}
