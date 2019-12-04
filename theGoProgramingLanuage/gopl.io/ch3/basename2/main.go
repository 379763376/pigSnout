// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 72.

// Basename2 reads file names from stdin and prints the base name of each one.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}

// basename removes directory components and 几种类型的print trailing .suffix.
// e.g., 几种类型的print => 几种类型的print, 几种类型的print.go => 几种类型的print, 几种类型的print/b/c.go => c, 几种类型的print/b.c.go => b.c
//!+
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//!-
