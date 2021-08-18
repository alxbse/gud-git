package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleNoSingleWord(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "single word",
			commit: object.Commit{
				Message: "Fix",
			},
			expected: false,
		},
		{
			name: "two words",
			commit: object.Commit{
				Message: "Fix test",
			},
			expected: true,
		},
		{
			name: "three words",
			commit: object.Commit{
				Message: "Fix test title",
			},
			expected: true,
		},
	}

	for _, tc := range tt {
		rule := RuleNoSingleWord{}
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
