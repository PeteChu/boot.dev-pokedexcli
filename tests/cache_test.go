package cache_test

import (
	"pokedexcli/internal/pokecache"
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := pokecache.NewCache(5 * time.Minute)
	defer cache.Stop()

	// Test cases
	testCases := []struct {
		name  string
		key   string
		value []byte
	}{
		{
			name:  "simple string",
			key:   "test-key",
			value: []byte("test-value"),
		},
		{
			name:  "JSON data",
			key:   "pokemon/pikachu",
			value: []byte(`{"name":"pikachu","type":"electric"}`),
		},
		{
			name:  "empty value",
			key:   "empty-key",
			value: []byte{},
		},
	}

	// Add values to cache
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache.Add(tc.key, tc.value)

			// Retrieve and verify
			got, found := cache.Get(tc.key)
			if !found {
				t.Errorf("expected key %s to be found in cache", tc.key)
			}

			if string(got) != string(tc.value) {
				t.Errorf("got %s, want %s", string(got), string(tc.value))
			}
		})
	}

	// Test non-existent key
	t.Run("non-existent key", func(t *testing.T) {
		_, found := cache.Get("non-existent-key")
		if found {
			t.Error("expected non-existent key to not be found")
		}
	})
}

func TestCacheExpiration(t *testing.T) {
	// Create cache with short expiration for testing
	interval := 10 * time.Millisecond
	cache := pokecache.NewCache(interval)
	defer cache.Stop()

	// Add test data
	testKey := "expiring-key"
	testValue := []byte("expiring-value")
	cache.Add(testKey, testValue)

	// Verify data is initially present
	_, found := cache.Get(testKey)
	if !found {
		t.Fatal("expected key to be found immediately after adding")
	}

	// Wait for expiration
	time.Sleep(interval * 2)

	// Verify data is gone after expiration
	_, found = cache.Get(testKey)
	if found {
		t.Error("expected key to be removed after expiration")
	}
}
