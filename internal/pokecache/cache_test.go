package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	cache.Add(
		"https://example.com",
		[]byte("testdata"),
	)

	val, ok := cache.Get("https://example.com")

	if !ok {
		t.Errorf("expected to find key")
		return
	}

	if string(val) != "testdata" {
		t.Errorf("expected testdata, got %s", string(val))
	}
}
