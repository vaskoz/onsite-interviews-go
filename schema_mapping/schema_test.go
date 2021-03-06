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
	{src: Schema{[]Field{
		{"country", []string{"nationality"}},
		{"state", []string{"st", "territory"}},
		{"city", []string{"cty", "burrow"}},
	}},
		target: Schema{[]Field{
			{"st", []string{"not country", "state", "sttt"}},
			{"burrow", []string{"hometown", "city"}},
			{"nationality", []string{"country", "nation", "residency"}},
		}},
		expected: map[string]string{
			"country": "nationality",
			"state":   "st",
			"city":    "burrow",
		}},
}

func TestMapping(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := Mapping(tc.src, tc.target); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkMapping(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Mapping(tc.src, tc.target)
		}
	}
}
