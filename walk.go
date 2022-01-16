// Copyright (c) 2022 Chris Palmer <chris@noncombatant.org>. All rights
// reserved. Use of this source code is governed by the Apache license that can
// be found in the LICENSE file.

package main

import (
	"os"
	"path"
)

func doWalk(root, directory string, pathnames chan Value) {
	if directory == root {
		defer close(pathnames)
	}

	entries, e := os.ReadDir(directory)
	if e != nil {
		pathnames <- Value{Error: e}
		return
	}

	for _, entry := range entries {
		name := path.Join(directory, entry.Name())
		if entry.IsDir() {
			doWalk(root, name, pathnames)
		} else {
			pathnames <- Value{String: name}
		}
	}
}

func Walk(root string) chan Value {
	pathnames := make(chan Value)
	go doWalk(root, root, pathnames)
	return pathnames
}
