package fib

import "testing"

func TestFibN(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		want  int
		want1 bool
	}{
		{"Negative number", -5, 0, false},
		{"Zero number", 0, 0, false},
		{"First number", 1, 1, true},
		{"Second number", 2, 1, true},
		{"Number 3", 3, 2, true},
		{"Number 10", 10, 55, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FibN(tt.n)
			if got != tt.want || got1 != tt.want1 {
				t.Errorf("FibN(%d) = (%d, %v), want (%d, %v)", tt.n, got, got1, tt.want, tt.want1)
			}
		})
	}
}
