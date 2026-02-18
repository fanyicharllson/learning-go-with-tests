package intergers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("add numbers", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5

		testLogs(t, sum, expected)

	})
}

func testLogs(t testing.TB, sum, expected int) {
	if sum != expected {
		t.Errorf("sum %d expected %d", sum, expected)
	}
}
