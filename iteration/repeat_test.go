package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("expects 5 a's character", func(t *testing.T) {
		repeat := Repeat("a")
		expected := "aaaaa"

		if repeat != expected {
			t.Errorf("repeat %q expected %q", repeat, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
