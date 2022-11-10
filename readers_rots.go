package main

import (
	"io"
	"os"
	"strings"
)

type rot13reader struct {
	r *strings.Reader
}

func (r rot13reader) Read(b []byte) (int, error) {
	x, err_v := r.r.Read(b)
	if err_v != nil {
		return 0, err_v
	}

	for i := 0; i < x; i++ {

		if b[i] > 'z'-13 {
			b[i] = b[i] - 13
		} else {
			b[i] = b[i] + 13
		}

	}
	// fmt.Println(x)

	return x, nil

}

func main() {

	s := strings.NewReader("Penpxvat plcuref vf sha!")
	// s := rot13reader{"Hello"}
	r := rot13reader{s}
	io.Copy(os.Stdout, r)
}

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// )

// func main() {
// 	r := strings.NewReader("Hello, Reader!")

// 	b := make([]byte, 8)
// 	for {
// 		n, err := r.Read(b)
// 		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
// 		fmt.Printf("b[:n] = %q\n", b[:n])
// 		if err == io.EOF {
// 			break
// 		}
// 	}
// }
