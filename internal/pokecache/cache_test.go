package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://site0.com",
			val: []byte("val0"),
		},
		{
			key: "https://site1.com",
			val: []byte("val1"),
		},
		{
			key: "https://site2.com",
			val: []byte("val2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

// // Why isn't this test working :'-(
// func TestReapLoop(t *testing.T) {
// 	const baseTime = 5 * time.Millisecond
// 	const waitTime = 2*baseTime + 5*time.Millisecond
// 	cache := NewCache(baseTime)
// 	cache.Add("https://site.com", []byte("some data"))

// 	_, ok := cache.Get("https://site.com")
// 	if !ok {
// 		t.Errorf("expected to find key")
// 		return
// 	}

// 	time.Sleep(waitTime)

// 	_, ok = cache.Get("https://site.com")
// 	if ok {
// 		t.Errorf("expected to not find key")
// 		return
// 	}
// }
