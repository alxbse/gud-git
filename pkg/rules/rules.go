package rules

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleSuggestion struct {
	Highlight  string
	Title      string
	Verb       string
	Suggestion string
	Branch     string
	Error      string
}

type Rule interface {
	Validate(commit *object.Commit) (bool, string)
	BreakWhenInvalid() bool
}

func HighlightRed(in string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", in)
}
