package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleNoConventionalSpec(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "valid title",
			commit: object.Commit{
				Message: "This is a valid title",
			},
			expected: true,
		},
		{
			name: "example invalid",
			commit: object.Commit{
				Message: "example: this is a conventional spec style title",
			},
			expected: false,
		},
		{
			name: "invalid title with scope",
			commit: object.Commit{
				Message: "example(test): whatever is in parenthesis is the scope",
			},
			expected: false,
		},
		{
			name: "invalid title with multi word style",
			commit: object.Commit{
				Message: "executor image: fix USER environment variable",
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		rule := RuleNoConventionalSpec{}
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
