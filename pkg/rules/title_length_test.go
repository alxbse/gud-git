package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleTitleLength(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "short title",
			commit: object.Commit{
				Message: "This is a short title",
			},
			expected: true,
		},
		{
			name: "long title",
			commit: object.Commit{
				Message: "This is a commit message with a title that has a length which is way too long",
			},
			expected: false,
		},
		{
			name: "special case for revert",
			commit: object.Commit{
				Message: "Revert 'This is a commit message with a title that has a length which is way too long'",
			},
			expected: true,
		},
	}

	for _, tc := range tt {
		rule := RuleTitleLength{}
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
