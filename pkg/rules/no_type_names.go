package rules

import (
	"strings"
	"text/template"
	"unicode"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func isUpperOrPunct(r rune) bool {
	if unicode.IsUpper(r) {
		return true
	}

	if unicode.IsPunct(r) {
		return true
	}

	return false
}

type RuleNoTypeNames struct{}

func (n *RuleNoTypeNames) BreakWhenInvalid() bool {
	return false
}

func (n *RuleNoTypeNames) Suggestion(ruleSuggestion *RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must not use internal details\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }} rewrite title to use natural language"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, ruleSuggestion)

	return b.String()
}

func (n *RuleNoTypeNames) Validate(c *object.Commit) (bool, string) {
	lines := strings.Split(c.Message, "\n")
	words := strings.Fields(lines[0])

	x := len(words[0])
	for _, word := range words[1:] {
		fields := strings.FieldsFunc(word, isUpperOrPunct)
		if len(fields) > 1 {
			color := HighlightRed(word)
			title := strings.Replace(lines[0], word, color, 1)
			highlight := strings.Repeat(" ", x+1) + "^"

			ruleSuggestion := RuleSuggestion{
				Title:     title,
				Highlight: highlight,
			}

			return false, n.Suggestion(&ruleSuggestion)
		}

		x = x + len(word+" ")
	}

	return true, ""
}
