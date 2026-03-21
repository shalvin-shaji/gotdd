package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Run test with character 'a'", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Run test with character 'c'", func(t *testing.T) {
		repeated := Repeat("c")
		expected := "ccccc"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
