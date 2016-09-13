//
// hamming distance calculations in Go
//
// https://github.com/steakknife/hamming
//
// Copyright © 2014, 2015, 2016 Barry Allard
//
// MIT license
//
package hamming

import (
	"fmt"
	"io/ioutil"
	"os"
)

func writeToTempFile(x int) {
	tmpfile, err := ioutil.TempFile("", "go-benchmark-dead-code-elimination-preventer")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())

	if n, err := fmt.Fprintf(tmpfile, "%d", x); err != nil {
		panic(err)
	} else if n == 0 {
		panic("couldn't write temp file")
	}

	if err := tmpfile.Close(); err != nil {
		panic(err)
	}
}
