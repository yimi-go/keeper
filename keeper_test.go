package keeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeeper_Get(t *testing.T) {
	count := 0
	k := NewKeeper(func(_ string) int {
		count++
		return count
	})
	v := k.Get("a")
	assert.Equal(t, 1, v)
	v = k.Get("b")
	assert.Equal(t, 2, v)
	v = k.Get("a")
	assert.Equal(t, 1, v)
	v = k.Get("b")
	assert.Equal(t, 2, v)
	assert.Equal(t, 2, count)
}

func TestKeeper_Reset(t *testing.T) {
	k := NewKeeper(func(key string) int {
		return 1
	}).(*keeper[string, int])
	k.Get("key")
	call := false
	k.Reset(func(key string) int {
		call = true
		return 1
	})

	length := 0
	for range k.store.Load().(map[string]int) {
		length++
	}
	assert.Equal(t, 0, length)

	k.Get("key")
	assert.True(t, call)

	defer func() {
		assert.NotNil(t, recover())
	}()
	k.Reset(nil)
}

func TestNewKeeper(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover())
	}()
	NewKeeper[string, int](nil)
}
