package rules

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleImperativeMood struct {
	Prepositions []string
}

func (r *RuleImperativeMood) BreakWhenInvalid() bool {
	return false
}

func (r *RuleImperativeMood) Validate(commit *object.Commit) (bool, string) {
	lines := strings.Split(commit.Message, "\n")
	words := strings.Fields(lines[0])
	if len(words) == 1 {
		return true, ""
	}

	for _, preposition := range r.Prepositions {
		if strings.ToLower(words[1]) == preposition {
			color := HighlightRed(words[1])
			title := strings.Replace(lines[0], words[1], color, 1)
			highlight := fmt.Sprint(strings.Repeat(" ", len(words[0])+1), strings.Repeat("^", len(words[1])))
			rs := RuleSuggestion{
				Title:     title,
				Highlight: highlight,
			}
			return false, r.Suggestion(rs)
		}
	}
	return true, ""
}

func (r *RuleImperativeMood) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must use imperative mood\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }} remove preposition after verb"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
