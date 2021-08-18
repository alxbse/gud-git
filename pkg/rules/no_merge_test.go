package rules

import (
	"reflect"
	"testing"
)

func TestRegexp(t *testing.T) {
	tt := []struct {
		name     string
		title    string
		expected []string
	}{
		{
			name:     "hello",
			title:    "Merge branch 'hello'",
			expected: []string{"hello"},
		},
		{
			name:     "remote",
			title:    "Merge remote-tracking branch 'hello'",
			expected: []string{"hello"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := re.FindStringSubmatch(tc.title)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %q, got %q", tc.expected, actual)
			}
		})
	}
}
