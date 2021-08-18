package rules

import (
	"regexp"
	"strings"
	"text/template"

	"github.com/go-git/go-git/v5/plumbing/object"
)

var (
	re = regexp.MustCompile("Merge (?:remote-tracking )?branch '(.*)'")
)

type RuleNoMerge struct {
}

func (r *RuleNoMerge) BreakWhenInvalid() bool {
	return true
}

func (r *RuleNoMerge) Validate(commit *object.Commit) (bool, string) {
	if commit.NumParents() > 1 {
		lines := strings.Split(commit.Message, "\n")
		rs := RuleSuggestion{}
		h := re.FindStringSubmatch(lines[0])
		if h != nil {
			rs.Branch = h[1]
		} else {
			rs.Branch = ""
		}
		return false, r.Suggestion(rs)
	}
	return true, ""
}

func (r *RuleNoMerge) Suggestion(rs RuleSuggestion) string {
	var b strings.Builder
	input := "\033[31merror\033[0m: merge commits are not allowed\n" +
		"details: Instead of '\033[31mgit merge {{ .Branch }}\033[0m' use 'git rebase {{ .Branch }}'"
	t := template.Must(template.New("hello").Parse(input))
	t.Execute(&b, rs)
	return b.String()
}
