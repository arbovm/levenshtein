package levenshtein

import (
	"fmt"
	"testing"
)

var distanceTests = []struct {
	first  string
	second string
	wanted int
}{
	{"a",    "a",  0},
	{"ab",  "ab",  0},
	{"ab",  "aa",  1},
	{"ab",  "aa",  1},
	{"ab", "aaa",  2},
	{"bbb",  "a",  3},
	{"kitten", "sitting", 3},
}

func TestDistance(t *testing.T) {

	for index, distanceTest := range distanceTests {
		result := Distance(distanceTest.first, distanceTest.second)
		if result != distanceTest.wanted {
			output := fmt.Sprintf("%v \t distance of %v and %v should be %v but was %v.",
				index, distanceTest.first, distanceTest.second, distanceTest.wanted, result)
			t.Errorf(output)
		}
	}
}
