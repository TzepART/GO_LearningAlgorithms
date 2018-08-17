package methods_and_interfaces

import (
	"io"
	"os"
	"strings"
)

type deRot13Reader struct {
	r io.Reader
}

type rot13Reader struct {
	r io.Reader
}

func deRot13byte(sb byte) byte {
	s := rune(sb)
	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		sb += 13
	}
	if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
		sb -= 13
	}

	return sb
}

func rot13byte(sb byte) byte {
	s := rune(sb)
	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		sb += 13
	}
	if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
		sb -= 13
	}

	return sb
}

func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13byte(p[i])
	}
	return
}

func (deRot13 deRot13Reader) Read(p []byte) (n int, err error) {
	n, err = deRot13.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = deRot13byte(p[i])
	}
	return
}

func ReadersExercise23() {
	s := strings.NewReader("You cracked the code!")
	r := deRot13Reader{s}
	io.Copy(os.Stdout, &r)

	deS := strings.NewReader("Lbh penpxrq gur pbqr!")
	deR := rot13Reader{deS}
	io.Copy(os.Stdout, &deR)
}
