package main

import (
	"math"
	"testing"
)

func floatCompare(a, b float64) bool {
	return math.Abs(a-b) < 0.001
}

func TestFloatCompare(t *testing.T) {
	for _, c := range []struct {
		name string
		a    float64
		b    float64
		want bool
	}{
		{
			name: "case1",
			a:    1.0,
			b:    1.0,
			want: true,
		},
		{
			name: "case2",
			a:    1.0,
			b:    1.000001,
			want: true,
		},
		{
			name: "case3",
			a:    1.0,
			b:    1.01,
			want: false,
			// want: true,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := floatCompare(c.a, c.b); got != c.want {
				t.Errorf("floatCompare(%v, %v) got %v, want %v", c.a, c.b, got, c.want)
			}
		})
	}
}
