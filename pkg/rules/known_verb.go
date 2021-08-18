package rules

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type ruleKnownVerbSuggestion struct {
	RuleSuggestion
	RuleKnownVerb
}

type RuleKnownVerb struct {
	KnownVerbs []string
}

func (r *RuleKnownVerb) BreakWhenInvalid() bool {
	return true
}

func (r *RuleKnownVerb) Validate(commit *object.Commit) (bool, string) {
	lines := strings.Split(commit.Message, "\n")
	words := strings.Fields(lines[0])

	for _, verb := range r.KnownVerbs {
		if strings.HasPrefix(strings.Title(words[0]), strings.Title(verb)) {
			return true, ""
		}
	}
	color := fmt.Sprint("\033[31m", words[0], "\033[0m")
	title := strings.Replace(lines[0], words[0], color, 1)
	rs := ruleKnownVerbSuggestion{
		RuleSuggestion{
			Title:     title,
			Highlight: strings.Repeat("^", len(words[0])),
		},
		RuleKnownVerb{
			KnownVerbs: r.KnownVerbs,
		},
	}
	return false, r.Suggestion(rs)
}

func (r *RuleKnownVerb) Suggestion(rs ruleKnownVerbSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must start with known verb\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }} replace with any of the known verbs:\n" +
		"{{ range $verb := .KnownVerbs }}           '{{ $verb }}'\n{{ end }}"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
