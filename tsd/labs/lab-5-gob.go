package main

import "bytes"
import "encoding/gob"
import "fmt"

// encode two regular values into a string
// that can be saved in a file.
func enc(x1 int, x2 string) string {
	w := new(bytes.Buffer)
	e := gob.NewEncoder(w)
  e.Encode(x1)
  e.Encode(x2)
	return string(w.Bytes())
}

// decode a string originally produced by enc() and
// return the original values.
func dec(buf string) (int, string) {
	r := bytes.NewBuffer([]byte(buf))
	d := gob.NewDecoder(r)
  var x1 int
  var x2 string
  d.Decode(&x1)
  d.Decode(&x2)
  return x1, x2
}

func main() {
  buf := enc(99, "hello")
  x1, x2 := dec(buf)
  fmt.Printf("%v %v\n", x1, x2)
}
