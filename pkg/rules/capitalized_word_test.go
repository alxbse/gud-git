package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleCapitalizedWord(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "valid title",
			commit: object.Commit{
				Message: "Start title with a capitalized verb",
			},
			expected: true,
		},
		{
			name: "invalid title",
			commit: object.Commit{
				Message: "start title with a lowercase verb",
			},
			expected: false,
		},
		{
			name: "excessive capitalization",
			commit: object.Commit{
				Message: "This Is too Much capitalization",
			},
			expected: false,
		},
		{
			name: "revert message is a special case",
			commit: object.Commit{
				Message: "Revert 'This is Too Much Capitalization'",
			},
			expected: true,
		},
		{
			name: "all caps wordss should pass",
			commit: object.Commit{
				Message: "Title with ALLCAPS words is OK",
			},
			expected: true,
		},
		{
			name: "version number should not count as capitalization",
			commit: object.Commit{
				Message: "Version like 1.2.3 is OK",
			},
			expected: true,
		},
	}

	for _, tc := range tt {
		rule := RuleCapitalizedWord{}
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
