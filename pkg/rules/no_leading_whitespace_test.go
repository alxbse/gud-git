package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleNoLeadingWhitespace(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "no leading whitespace",
			commit: object.Commit{
				Message: "No leading whitespace",
			},
			expected: true,
		},
		{
			name: "leading space",
			commit: object.Commit{
				Message: " Leading space",
			},
			expected: false,
		},
		{
			name: "leading tab",
			commit: object.Commit{
				Message: "	Leading tab",
			},
			expected: false,
		},
	}

	rule := RuleNoLeadingWhitespace{}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
