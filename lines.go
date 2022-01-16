// Copyright (c) 2022 Chris Palmer <chris@noncombatant.org>. All rights
// reserved. Use of this source code is governed by the Apache license that can
// be found in the LICENSE file.

package main

import (
	"bufio"
	"io"
)

func doLines(reader io.Reader, lines chan Value) {
	defer close(lines)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines <- Value{String: scanner.Text()}
	}
	if e := scanner.Err(); e != nil {
		lines <- Value{Error: e}
	}
}

func Lines(reader io.Reader) chan Value {
	lines := make(chan Value)
	go doLines(reader, lines)
	return lines
}
