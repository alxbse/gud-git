package rules

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleNoFixup struct {
}

func (r *RuleNoFixup) BreakWhenInvalid() bool {
	return true
}

func (r *RuleNoFixup) Validate(commit *object.Commit) (bool, string) {
	if strings.HasPrefix(commit.Message, "fixup!") {
		lines := strings.Split(commit.Message, "\n")
		words := strings.Fields(lines[0])
		color := fmt.Sprintf("\033[31m%s\033[0m", words[0])
		title := strings.Replace(lines[0], words[0], color, 1)
		rs := RuleSuggestion{
			Title:     title,
			Highlight: strings.Repeat("^", len(words[0])),
		}
		return false, r.Suggestion(rs)
	}
	return true, ""
}

func (r *RuleNoFixup) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: fixup commits must be squashed\n" +
		"details: '\033[31m{{ .Title }}\033[0m'\n" +
		"          {{ .Highlight }} use 'git rebase --interactive --autosquash' to squash"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
