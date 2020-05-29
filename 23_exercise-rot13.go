package main

import (
	"io"
	"os"
	"strings"
	// "fmt"
)

type rot13Reader struct {
	r io.Reader
}

// this infinite stream, don't use range
func (rot *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)
	// fmt.Printf("n = %v, err = %v, b = %v\n", n, err, n)
	// fmt.Printf("b[:n] = %q\n", b[:n])

	for i:=0; i<n; i++ {
		b[i] = rot13Chiper(b[i])
		// fmt.Println(b[i])
	}

	return n, err
}

func rot13Chiper(in byte) (out byte) {
	out = in
	if in<91 && in>64 {
		// capital alphabet
		if in<78 {
			out = in+13
		} else if in<91 {
			out = in-13
		}
	} else if in<123 && in>96 {
		// small alphabet
		if in<110 {
			out = in+13
		} else if  in>96 {
			out = in-13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}