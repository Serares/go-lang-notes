package methods_and_interfaces

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (the_reader *rot13Reader) Read(b []byte) (int, error) {
	n, err := the_reader.r.Read(b)
	decrypted_output := make([]byte, n)
	if err != nil {
		return 0, err
	}
	the_string := b[:n]

	for i, ch := range the_string {
		if ch >= 'A' && ch <= 'Z' {
			decrypted_output[i] = 'A' + (ch-'A'+13)%26
		} else if ch >= 'a' && ch <= 'z' {
			decrypted_output[i] = 'a' + (ch-'a'+13)%26
		}
	}
	fmt.Println("The decrypted output:", string(decrypted_output))
	return len(decrypted_output), nil
}

func UseTheReader() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
