package skyline

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	buildings []Building
	expected  []Point
}{
	{[]Building{
		{5, 7, 10},
		{8, 20, 5},
		{9, 12, 7},
		{10, 15, 12},
		{18, 25, 6},
	},
		[]Point{{5, 0}, {5, 10}, {7, 10}, {7, 0}, {8, 0}, {8, 5}, {9, 5}, {9, 7}, {10, 7}, {10, 12},
			{15, 12}, {15, 5}, {18, 5}, {18, 6}, {25, 6}, {25, 0}}},
}

func TestSkyline(t *testing.T) {
	for _, tc := range testcases {
		if skyline := Skyline(tc.buildings); !reflect.DeepEqual(skyline, tc.expected) {
			t.Errorf("Expected %v got %v", tc.expected, skyline)
		}
	}
}

func BenchmarkSkyline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Skyline(tc.buildings)
		}
	}
}
