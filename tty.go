package tty

import (
	"fmt"
)

const (
	leftTop, top, rightTop          = '\u250C', '\u2500', '\u2510'
	left, inner, right              = '\u2502', '\u0020', '\u2502'
	leftBottom, bottom, rightBottom = '\u2514', '\u2500', '\u2518'
)

type Color int

const (
	FgBlack   = Color(30)
	FgRed     = Color(31)
	FgGreen   = Color(32)
	FgYellow  = Color(33)
	FgBlue    = Color(34)
	FgMagenta = Color(35)
	FgCyan    = Color(36)
	FgWhite   = Color(37)

	BgBlack   = Color(40)
	BgRed     = Color(41)
	BgGreen   = Color(42)
	BgYellow  = Color(43)
	BgBlue    = Color(44)
	BgMagenta = Color(45)
	BgCyan    = Color(46)
	BgWhite   = Color(47)
)

type Align int

const (
	AlignLeft   = Align(0)
	AlignRight  = Align(1)
	AlignCenter = Align(2)
)

type Attr int

const (
	AttrNormal     = Attr(0)
	AttrBright     = Attr(1)
	AttrDim        = Attr(2)
	AttrUnderscore = Attr(4)
	AttrBlink      = Attr(5)
	AttrReverse    = Attr(7)
	AttrHidden     = Attr(8)
)

var (
	spinner1 = []int{'|', '/', '-', '\\'}
)

func Spinner(progress <-chan int) {
	for i := range progress {
		fmt.Printf("\x08%c", spinner1[i%4])
	}
}

func Progressbar(width int, progress <-chan int) {

	// background
	for i := 0; i < width; i++ {
		fmt.Print("\u2591")
	}
	CursorLeft(width)

	// progress indicator
	max := <-progress
	step := float64(width) / float64(max)

	last := width
	for i := range progress {
		position := int(float64(i) * step)

		if position >= last || position == 0 {
			continue
		}

		for j := last; j > position; j-- {
			fmt.Print("\u258B")
		}
		last = position
	}
}

func CursorUp(num int) {
	fmt.Printf("\x1b[%dA", num)
}

func CursorDown(num int) {
	fmt.Printf("\x1b[%dB", num)
}

func CursorLeft(num int) {
	fmt.Printf("\x1b[%dD", num)
}

func CursorRight(num int) {
	fmt.Printf("\x1b[%dC", num)
}

func CursorPosition(col, row int) {
	fmt.Printf("\x1b[%d;%df", row, col)
}

func Colors(fg, bg Color) {
	fmt.Printf("\x1b[%d;%dm", fg, bg)
}

func Attribute(attr Attr) {
	fmt.Printf("\x1b[%dm", attr)
}

func Text(width int, align Align, text string) {

	// clear
	for c := 0; c < width-2; c++ {
		fmt.Printf("%c", inner)
	}
	CursorLeft(width)

	// start
	start := 0
	switch align {
	case AlignLeft:
		break
	case AlignRight:
		start = width - len(text)
		break
	case AlignCenter:
		start = (width - len(text)) / 2
		break
	}
	CursorRight(start + 2)

	// text
	fmt.Print(text)
}

func Shell(width, height int) {

	// top
	fmt.Printf("%c", leftTop)
	for c := 0; c < width-2; c++ {
		fmt.Printf("%c", top)
	}
	fmt.Printf("%c", rightTop)

	// inner
	for r := 0; r < height-2; r++ {
		CursorLeft(width)
		CursorDown(1)
		fmt.Printf("%c", left)
		for c := 0; c < width-2; c++ {
			fmt.Printf("%c", inner)
		}
		fmt.Printf("%c", right)
	}

	// bottom
	CursorLeft(width)
	CursorDown(1)
	fmt.Printf("%c", leftBottom)
	for c := 0; c < width-2; c++ {
		fmt.Printf("%c", bottom)
	}
	fmt.Printf("%c", rightBottom)
}

func Reset() {
	fmt.Print("\x1bc")
}

func ResetAttributes() {
	Attribute(AttrNormal)
}

func HideCursor() {
	fmt.Print("\x1b[?25l")
}

func ShowCursor() {
	fmt.Print("\x1b[?25h")
}

func SaveCursor() {
	fmt.Print("\x1b[s")
}

func RestoreCursor() {
	fmt.Print("\x1b[u")
}
