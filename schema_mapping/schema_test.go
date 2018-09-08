package schema

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	src, target Schema
	expected    map[string]string
}{
	{src: Schema{[]Field{{"country", []string{"cnt", "countryLong"}}}},
		target:   Schema{[]Field{{"countryLong", []string{"country"}}}},
		expected: map[string]string{"country": "countryLong"}},
}

func TestMapping(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Mapping(tc.src, tc.target); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}
