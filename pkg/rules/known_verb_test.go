package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleKnownVerb(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "test known verb",
			commit: object.Commit{
				Message: "add feature",
			},
			expected: true,
		},
		{
			name: "test unknown verb",
			commit: object.Commit{
				Message: "recalculate numbers",
			},
			expected: false,
		},
		{
			name: "test mixed case",
			commit: object.Commit{
				Message: "Add feature",
			},
			expected: true,
		},
		{
			name: "test present form",
			commit: object.Commit{
				Message: "adding feature",
			},
			expected: true,
		},
	}

	rule := RuleKnownVerb{
		KnownVerbs: []string{
			"add",
			"update",
			"delete",
			"remove",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t\n", tc.expected, actual)
			}
		})
	}
}
