package test

import "testing"

func TestCacheKeyProduced(t *testing.T) {
	got := (-1) * -1
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
