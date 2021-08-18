package rules

import (
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleNoConventionalSpec struct {
}

func (r *RuleNoConventionalSpec) BreakWhenInvalid() bool {
	return true
}

func (r *RuleNoConventionalSpec) Validate(commit *object.Commit) (bool, string) {
	lines := strings.Split(commit.Message, "\n")
	words := strings.Fields(lines[0])
	repeat := 0
	for i, word := range words {
		repeat = repeat + len(word)
		if strings.HasSuffix(word, ":") {
			color := HighlightRed(word)
			title := strings.Replace(lines[0], word, color, 1)
			rs := RuleSuggestion{
				Title:     title,
				Highlight: strings.Repeat("^", repeat+i),
			}
			return false, r.Suggestion(rs)
		}
	}
	return true, ""
}

func (r *RuleNoConventionalSpec) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must not use conventional syntax\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }} rewrite title to use standard syntax"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
