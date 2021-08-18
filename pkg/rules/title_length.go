package rules

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

const (
	DefaultTitleMaxLength = 72
)

type RuleTitleLength struct {
	MaxLength int
}

type lengthRule struct {
	RuleSuggestion
	RuleTitleLength
}

func (r *RuleTitleLength) BreakWhenInvalid() bool {
	return false
}

func (r *RuleTitleLength) Validate(commit *object.Commit) (bool, string) {
	if r.MaxLength == 0 {
		r.MaxLength = DefaultTitleMaxLength
	}
	s := strings.Split(commit.Message, "\n")
	fields := strings.Fields(s[0])

	// Since `git revert` always act on already merged commits title does not matter.
	if fields[0] == "Revert" {
		return true, ""
	}

	if len(s[0]) > r.MaxLength {
		rt := lengthRule{
			RuleSuggestion: RuleSuggestion{
				Title: s[0],
			},
			RuleTitleLength: RuleTitleLength{
				MaxLength: r.MaxLength,
			},
		}
		return false, r.Suggestion(rt)
	}
	return true, ""
}

func (r *RuleTitleLength) Suggestion(rt lengthRule) string {
	prefix := rt.Title[0:r.MaxLength]
	suffix := rt.Title[r.MaxLength:]
	title := fmt.Sprint(prefix, "\033[31m", suffix, "\033[0m")
	highlight := fmt.Sprint(strings.Repeat(" ", len(prefix)), strings.Repeat("^", len(suffix)))

	rt.Title = title
	rt.Highlight = highlight

	var b strings.Builder
	input := "\033[31merror\033[0m: title must be shorter than {{ .MaxLength }} characters\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }}"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rt)
	return b.String()
}
