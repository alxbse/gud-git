package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleImperativeMoodSpec(t *testing.T) {
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
			name: "Second word is of",
			commit: object.Commit{
				Message: "Update of various things",
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		rule := RuleImperativeMood{
			Prepositions: []string{"of"},
		}
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
