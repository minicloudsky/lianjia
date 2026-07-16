package biz

import (
	"math/rand"
	"testing"

	"bou.ke/monkey"
)

func TestRandomUA_WithMonkeyPatch(t *testing.T) {
	t.Run("mock rand.Intn", func(t *testing.T) {
		// Monkey patch rand.Intn to always return index 2
		guard := monkey.Patch(rand.Intn, func(n int) int {
			return 2
		})
		defer guard.Restore()

		ua := RandomUA()
		expectedUA := string(userAgents[2])
		if ua != expectedUA {
			t.Errorf("expected UA %q, got %q", expectedUA, ua)
		}
	})

	t.Run("mock RandomUA itself", func(t *testing.T) {
		// Monkey patch RandomUA itself to return a completely custom user agent
		customUA := "Mozilla/5.0 (CustomUA) Test/1.0"
		guard := monkey.Patch(RandomUA, func() string {
			return customUA
		})
		defer guard.Restore()

		ua := RandomUA()
		if ua != customUA {
			t.Errorf("expected UA %q, got %q", customUA, ua)
		}
	})
}
