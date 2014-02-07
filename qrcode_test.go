package qrcode

import (
	"testing"
)

func TestVersionDeterminesSize(t *testing.T) {
	table := []struct {
		version  int
		expected int
	}{
		{1, 21},
		{2, 25},
		{6, 41},
		{7, 45},
		{14, 73},
		{40, 177},
	}

	for _, test := range table {
		size := Version(test.version).PatternSize()
		if size != test.expected {
			t.Errorf("Version %2d, expected %3d but got %3d",
				test.version, test.expected, size)
		}
	}
}
