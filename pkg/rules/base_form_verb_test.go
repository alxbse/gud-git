package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleBaseFormVerb(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "valid base form",
			commit: object.Commit{
				Message: "Add new feature",
			},
			expected: true,
		},
		{
			name: "invalid past form",
			commit: object.Commit{
				Message: "Added new feature",
			},
			expected: false,
		},
		{
			name: "invalid singular present form",
			commit: object.Commit{
				Message: "Adds new feature",
			},
			expected: false,
		},
		{
			name: "invalid present form",
			commit: object.Commit{
				Message: "Adding new feature",
			},
			expected: false,
		},
	}

	for _, tc := range tt {
		rule := RuleBaseFormVerb{}
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
