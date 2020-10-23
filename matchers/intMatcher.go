package matchers

import (
	"github.com/matheus0214/book/testes/search"
	"regexp"
)

// Register IntMacther in init program
func init() {
	var matcher IntMatcher
	search.Register("Int", matcher)
}

type IntMatcher struct{}

// IntMatchers - search for integer values in file
func (iM IntMatcher) Search(values []*search.Values) ([]*search.Result, error) {
	var r []*search.Result

	for _, value := range values {
		verify, _ := regexp.Match(`^(\d+)$`, []byte(value.Content))
		if verify {
			result := &search.Result{Type: "Int", Content: value.Content}
			r = append(r, result)
		}
	}

	return r, nil
}
