package prime

import (
	"fmt"
	"testing"
)

func TestPrimeTableDriven(t *testing.T) {
	var tests = []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{179, true},
		{401, true},
	}

	for _, tt := range tests {
		tname := fmt.Sprintf("Prime is %d", tt.n)
		t.Run(tname, func(t *testing.T) {
			ans := Prime(tt.n)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})

	}
}

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prime(1000000007)
	}
}
