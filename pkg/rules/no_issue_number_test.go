package rules

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestRuleNoIssueNumber(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "no issue number",
			commit: object.Commit{
				Message: "This is a title without any issue numbers",
			},
			expected: true,
		},
		{
			name: "jira ticket at the start",
			commit: object.Commit{
				Message: "JRA-12345 reference to jira ticket at the start",
			},
			expected: false,
		},
		{
			name: "jira ticket in the middle",
			commit: object.Commit{
				Message: "Refence to JRA-12345 in the middle",
			},
			expected: false,
		},
		{
			name: "github issue at the end",
			commit: object.Commit{
				Message: "Reference to issue #1234",
			},
			expected: false,
		},
		{
			name: "cve identifier at the end",
			commit: object.Commit{
				Message: "Reference to vulnerability CVE-2021-1234567",
			},
			expected: false,
		},
		{
			name: "short id is good",
			commit: object.Commit{
				Message: "Fix something in K2",
			},
			expected: true,
		},
	}

	for _, tc := range tt {
		rule := RuleNoIssueNumber{}
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := rule.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
