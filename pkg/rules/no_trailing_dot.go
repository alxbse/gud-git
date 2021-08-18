package rules

import (
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleNoTrailingPunct struct {
}

func (r *RuleNoTrailingPunct) BreakWhenInvalid() bool {
	return false
}

func (r *RuleNoTrailingPunct) Validate(commit *object.Commit) (bool, string) {
	// Special case for 'git revert'
	if strings.HasPrefix(commit.Message, "Revert") {
		return true, ""
	}

	lines := strings.Split(commit.Message, "\n")
	title := lines[0]
	last, _ := utf8.DecodeLastRuneInString(title)
	if unicode.IsPunct(last) {
		rs := RuleSuggestion{
			Highlight:  strings.Repeat(" ", len(title)-1),
			Title:      title,
			Suggestion: title[len(title)-1:],
		}
		return false, r.Suggestion(rs)
	}
	return true, ""
}

func (r *RuleNoTrailingPunct) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must end with whitespace\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }}^ remove the '{{ .Suggestion }}'"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
