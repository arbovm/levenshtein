// Copyright (c) 2015, Arbo von Monkiewitsch All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package levenshtein

// The Levenshtein distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
// http://en.wikipedia.org/wiki/Levenshtein_distance
//
// This implemention is optimized to use O(min(m,n)) space.
// It is based on the optimized C version found here:
// http://en.wikibooks.org/wiki/Algorithm_implementation/Strings/Levenshtein_distance#C
func Distance(str1, str2 string) int {
	s1, s2 := []rune(str1), []rune(str2)
	lenS1, lenS2 := len(s1), len(s2)
	column := make([]int, lenS1+1)

	for i := range column {
		column[i] = i // initialise, in case lenS2 is empty
	}
	column[0] = lenS2 // in case str1 is empty

	for x := 0; x < lenS2; x++ {
		lastdiag := x
		for y := 0; y < lenS1; y++ {
			cost := 0
			if s1[y] != s2[x] {
				cost = 1
			}
			lastdiag, column[y+1] = column[y+1], min(column[y+1]+1, column[y]+1, lastdiag+cost)
		}
	}
	return column[lenS1]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c && b < a {
		return b
	}
	return c
}
