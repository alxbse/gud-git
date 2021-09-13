package rules

import (
	"fmt"
	"strings"
	"text/template"
	"unicode"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleCapitalizedWord struct {
}

func (r *RuleCapitalizedWord) BreakWhenInvalid() bool {
	return false
}

func (r *RuleCapitalizedWord) Validate(commit *object.Commit) (bool, string) {
	lines := strings.Split(commit.Message, "\n")
	words := strings.Fields(lines[0])

	// Since `git revert` always act on already merged commits title does not matter.
	if words[0] == "Revert" {
		return true, ""
	}

	if words[0] != strings.Title(words[0]) {
		color := fmt.Sprintf("\033[31m%s\033[0m", words[0])
		title := strings.Replace(lines[0], words[0], color, 1)
		rs := RuleSuggestion{
			Error:      "title must start with a capitalized word",
			Title:      title,
			Suggestion: fmt.Sprintf("use capitalized '%s' instead", strings.Title(words[0])),
			Highlight:  strings.Repeat("^", len(words[0])),
		}
		return false, r.Suggestion(rs)
	}

	highlights := []string{}
	highlights = append(highlights, strings.Repeat(" ", len(words[0])))
	for _, word := range words[1:] {
		isNumber := unicode.IsNumber(rune(word[0]))
		lower := strings.ToLower(word)
		if !isNumber && word == strings.Title(lower) {
			highlight := fmt.Sprint("^", strings.Repeat(" ", len(word)-1))
			highlights = append(highlights, highlight)
		} else {
			highlights = append(highlights, strings.Repeat(" ", len(word)))
		}
	}

	highlight := strings.Join(highlights, " ")
	if strings.Contains(highlight, "^") {
		highlight = strings.TrimRightFunc(highlight, unicode.IsSpace)
		rs := RuleSuggestion{
			Error:      "title must not use excessive capitalization",
			Title:      lines[0],
			Highlight:  highlight,
			Suggestion: "only capitalize the first word",
		}
		return false, r.Suggestion(rs)
	}
	return true, ""
}

func (r *RuleCapitalizedWord) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: {{ .Error }}\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }} {{ .Suggestion }}"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
