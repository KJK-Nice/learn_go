package main


import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
    n, err := reader.r.Read(b)
    if err == io.EOF {
        return 0, err
    }

    for i := range b {
        if b[i] >= 65 && b[i] <=90  {
            if b[i] >= 65 && b[i] < 77 {
                b[i] = b[i] + 13
            } else {
                b[i] = b[i] - 13
            }
        }
        if b[i] >= 97 && b[i] <= 122 {
            if b[i] >= 97 && b[i] < 109 {
                b[i] = b[i] + 13
            } else {
                b[i] = b[i] - 13
            }
        }
    }
    return n, nil
}

func main() {
    s := strings.NewReader("Lbh penpxrqgur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
