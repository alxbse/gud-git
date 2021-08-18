package rules

import (
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleNoSingleWord struct {
}

func (r *RuleNoSingleWord) BreakWhenInvalid() bool {
	return true
}

func (r *RuleNoSingleWord) Validate(commit *object.Commit) (bool, string) {
	lines := strings.Split(commit.Message, "\n")
	words := strings.Fields(lines[0])
	if len(words) == 1 {
		rs := RuleSuggestion{
			Title:     lines[0],
			Highlight: strings.Repeat("^", len(lines[0])),
		}
		return false, r.Suggestion(rs)
	}
	return true, ""
}

func (r *RuleNoSingleWord) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must be descriptive\n" +
		"details: '\033[31m{{ .Title }}\033[0m'\n" +
		"          {{ .Highlight }} add one or more words to make title descriptive"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
