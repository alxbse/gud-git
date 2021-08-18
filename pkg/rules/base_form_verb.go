package rules

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleBaseFormVerb struct {
}

func (r *RuleBaseFormVerb) BreakWhenInvalid() bool {
	return true
}

func (r *RuleBaseFormVerb) Validate(commit *object.Commit) (bool, string) {
	lines := strings.Split(commit.Message, "\n")
	words := strings.Fields(lines[0])
	suffixes := []string{
		"ing",
		"es",
		"ed",
		"s",
		"al",
	}
	for _, suffix := range suffixes {
		if strings.HasSuffix(words[0], suffix) {
			color := fmt.Sprintf("\033[31m%s\033[0m", words[0])
			title := strings.Replace(lines[0], words[0], color, 1)
			rs := RuleSuggestion{
				Highlight:  strings.Repeat("^", len(words[0])),
				Title:      title,
				Verb:       words[0],
				Suggestion: strings.Title(strings.TrimSuffix(words[0], suffix)),
			}
			suggestion := r.Suggestion(rs)
			return false, suggestion
		}
	}
	return true, ""
}

func (r *RuleBaseFormVerb) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must use base form verb.\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }} use base form '{{ .Suggestion }}' instead"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
