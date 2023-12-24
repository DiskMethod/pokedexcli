package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestSetGet(t *testing.T)  {
	const interval = 5 * time.Second
	type KeyValue struct {
		key string
		val []byte
	}
	cases := make([]KeyValue, 2)
	cases[0] = KeyValue{
		key: "https://example.com",
		val: []byte("testdata"),
	}
	cases[1] = KeyValue{
		key: "https://example.com/path",
		val: []byte("moretestdata"),
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Set(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Set("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
	}
}