package uuid

import (
	"testing"
)

func TestIntToUUID(t *testing.T) {
	for i := 0; i < 10000000; i++ {
		s := IntToUUID(i)
		if len(s) != 36 {
			t.Errorf("seed %d: wrong uuid length: %d, expect 36", i, len(s))
		}
		if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
			t.Errorf("seed %d: wrong uuid format: %s", i, s)
		}
	}
}

func BenchmarkIntToUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IntToUUID(i)
	}
}
