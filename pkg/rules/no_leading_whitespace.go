package rules

import (
	"strings"
	"unicode"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleNoLeadingWhitespace struct {
}

func (r *RuleNoLeadingWhitespace) BreakWhenInvalid() bool {
	return true
}

func (r *RuleNoLeadingWhitespace) Validate(commit *object.Commit) (bool, string) {
	trim := strings.TrimLeftFunc(commit.Message, unicode.IsSpace)
	if trim != commit.Message {
		return false, "Title must not contain leading whitespace"
	}
	return true, ""
}
