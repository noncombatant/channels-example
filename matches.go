// Copyright (c) 2022 Chris Palmer <chris@noncombatant.org>. All rights
// reserved. Use of this source code is governed by the Apache license that can
// be found in the LICENSE file.

package main

import (
	"regexp"
)

type Patterns []*regexp.Regexp

func matchAny(line string, patterns Patterns) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(line) {
			return true
		}
	}
	return false
}

func doMatches(lines chan Value, patterns Patterns, matches chan Value) {
	defer close(matches)

	for line := range lines {
		if line.Error != nil {
			matches <- Value{Error: line.Error}
		}
		if line.String != "" {
			if matchAny(line.String, patterns) {
				matches <- Value{String: line.String}
			}
		}
	}
}

func Matches(lines chan Value, patterns Patterns) chan Value {
	matches := make(chan Value)
	go doMatches(lines, patterns, matches)
	return matches
}

func MustCompilePatterns(patterns []string) Patterns {
	var results Patterns
	for _, p := range patterns {
		results = append(results, regexp.MustCompile(p))
	}
	return results
}
