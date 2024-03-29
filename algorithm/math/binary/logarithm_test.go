package binary

import "testing"

func TestLogBase2(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want uint32
	}{
		{"log2(1) = 0", 1, 0},
		{"log2(2) = 1", 2, 1},
		{"log2(4) = 2", 4, 2},
		{"log2(8) = 3", 8, 3},
		{"log2(16) = 4", 16, 4},
		{"log2(32) = 5", 32, 5},
		{"log2(64) = 6", 64, 6},
		{"log2(128) = 7", 128, 7},
		{"log2(256) = 8", 256, 8},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := LogBase2(test.n); got != test.want {
				t.Errorf("LogBase2() = %v, want %v", got, test.want)
			}
		})
	}
}
