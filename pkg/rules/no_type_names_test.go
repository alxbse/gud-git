package rules_test

import (
	"testing"

	"github.com/alxbse/gud-git/pkg/rules"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestNoTypeNamesRule(t *testing.T) {
	tt := []struct {
		name     string
		commit   object.Commit
		expected bool
	}{
		{
			name: "using separate words",
			commit: object.Commit{
				Message: "Update cluster template type",
			},
			expected: true,
		},
		{
			name: "using snake casing",
			commit: object.Commit{
				Message: "Update clusterTemplate type",
			},
			expected: false,
		},
		{
			name: "using pascal casing",
			commit: object.Commit{
				Message: "Update ClusterTemplate type",
			},
		},
		{
			name: "using kebab case",
			commit: object.Commit{
				Message: "Update cluster-template type",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			r := rules.RuleNoTypeNames{}

			actual, _ := r.Validate(&tc.commit)
			if actual != tc.expected {
				t.Errorf("expected %t, got %t", tc.expected, actual)
			}
		})
	}
}
