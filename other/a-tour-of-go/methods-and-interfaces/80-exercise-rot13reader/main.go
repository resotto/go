package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(buf []byte) (int, error) {
	size, err := r.r.Read(buf)
	if err == io.EOF {
		return size, err
	}

	for i, char := range buf {
		switch {
		case 'A' <= char && char <= 'M':
			buf[i] = 'N' + (char - 'A')
		case 'N' <= char && char <= 'Z':
			buf[i] = 'A' + (char - 'N')
		case 'a' <= char && char <= 'm':
			buf[i] = 'n' + (char - 'a')
		case 'n' <= char && char <= 'z':
			buf[i] = 'a' + (char - 'n')
		default:
			continue
		}
	}
	return size, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
