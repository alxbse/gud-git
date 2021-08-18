package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleNoTrailingPunct(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "no trailing punctuation",
			commit: object.Commit{
				Message: "Add a test",
			},
			expected: true,
		},
		{
			name: "trailing dot",
			commit: object.Commit{
				Message: "Add a test.",
			},
			expected: false,
		},
		{
			name: "trailing dash",
			commit: object.Commit{
				Message: "Add a test-",
			},
			expected: false,
		},
		{
			name: "trailing comma",
			commit: object.Commit{
				Message: "Add a test-",
			},
			expected: false,
		},
		{
			name: "special case for revert",
			commit: object.Commit{
				Message: "Revert 'Add a test'",
			},
			expected: true,
		},
	}

	for _, tc := range tt {
		rule := RuleNoTrailingPunct{}
		t.Run(tc.name, func(t *testing.T) {

			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
