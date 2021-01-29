package auth

import "testing"

func TestSameStrings(t *testing.T) {
	t.Run("when given different strings, returns `false`", func(t *testing.T) {
		want := false

		s1 := " Test"
		s2 := "Test"

		if got := SameStrings(s1, s2); got != want {
			t.Errorf("SameStrings() = %t, want %t", got, want)
		}
	})

	t.Run("when given the same strings, returns `true`", func(t *testing.T) {
		want := true

		s1 := "Test"
		s2 := "Test"

		if got := SameStrings(s1, s2); got != want {
			t.Errorf("SameStrings() = %t, want %t", got, want)
		}
	})
}

// TODO: add test for AuthMiddleware
