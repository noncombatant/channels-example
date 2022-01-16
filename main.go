// Copyright (c) 2022 Chris Palmer <chris@noncombatant.org>. All rights
// reserved. Use of this source code is governed by the Apache license that can
// be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

type Value struct {
	String string
	Error  error
}

func PrintValue(value Value) {
	if value.String != "" {
		fmt.Println(value.String)
	}
	if value.Error != nil {
		fmt.Fprintln(os.Stderr, value.Error)
	}
}

func main() {
	flag.Parse()
	subprogram := flag.Args()[0]
	arguments := flag.Args()[1:]

	if subprogram == "walk" {
		if len(arguments) == 0 {
			arguments = append(arguments, ".")
		}
		for _, pathname := range arguments {
			for filename := range Walk(pathname) {
				PrintValue(filename)
			}
		}
	} else if subprogram == "lines" {
		if len(arguments) == 0 {
			for line := range Lines(os.Stdin) {
				PrintValue(line)
			}
		} else {
			for _, pathname := range arguments {
				file, e := os.Open(pathname)
				if e != nil {
					fmt.Fprintln(os.Stderr, e)
					continue
				}
				for line := range Lines(file) {
					PrintValue(line)
				}
				file.Close()
			}
		}
	} else if subprogram == "matches" {
		patterns := MustCompilePatterns([]string{arguments[0]})
		if len(arguments) == 1 {
			for line := range Matches(Lines(os.Stdin), patterns) {
				PrintValue(line)
			}
		} else {
			for _, pathname := range arguments[1:] {
				file, e := os.Open(pathname)
				if e != nil {
					fmt.Fprintln(os.Stderr, e)
					continue
				}
				for line := range Matches(Lines(file), patterns) {
					PrintValue(line)
				}
				file.Close()
			}
		}
	}
}
