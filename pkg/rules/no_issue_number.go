package rules

import (
	"fmt"
	"regexp"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type RuleNoIssueNumber struct {
}

func (r *RuleNoIssueNumber) BreakWhenInvalid() bool {
	return false
}

func (r *RuleNoIssueNumber) Validate(commit *object.Commit) (bool, string) {
	re := regexp.MustCompile(`[A-Z#]+\-?[0-9]{2,}\-?[0-9]*`)
	lines := strings.Split(commit.Message, "\n")
	m := re.FindStringSubmatch(lines[0])
	if len(m) > 0 {
		color := fmt.Sprint("\033[31m", m[0], "\033[0m")
		title := strings.ReplaceAll(lines[0], m[0], color)
		indices := re.FindStringSubmatchIndex(lines[0])
		highlight := fmt.Sprint(strings.Repeat(" ", indices[0]), strings.Repeat("^", len(m[0])))
		rs := RuleSuggestion{
			Title:     title,
			Highlight: highlight,
		}
		return false, r.Suggestion(rs)
	}
	return true, ""
}

func (r *RuleNoIssueNumber) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: title must be descriptive\n" +
		"details: '{{ .Title }}'\n" +
		"          {{ .Highlight }} move issue number to body"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
