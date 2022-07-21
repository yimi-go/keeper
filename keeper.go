package keeper

import "sync/atomic"

type Keeper[K comparable, V any] interface {
	Get(key K) V
	Reset(new func(key K) V)
	Clear()
}

type keeper[K comparable, V any] struct {
	new   func(key K) V
	store atomic.Value
}

func (k *keeper[K, V]) Get(key K) V {
	m := k.store.Load().(map[K]V)
	if v, ok := m[key]; ok {
		return v
	}
	nm := make(map[K]V, len(m)+1)
	for k, v := range m {
		nm[k] = v
	}
	v := k.new(key)
	nm[key] = v
	k.store.Store(nm)
	return v
}

func (k *keeper[K, V]) Reset(new func(key K) V) {
	if new == nil {
		panic("keeper: can not assign nil as the new function")
	}
	k.new = new
	k.Clear()
}

func (k *keeper[K, V]) Clear() {
	k.store.Store(map[K]V{})
}

func NewKeeper[K comparable, V any](new func(key K) V) Keeper[K, V] {
	if new == nil {
		panic("keeper: can not assign nil as the new function")
	}
	k := &keeper[K, V]{
		new: new,
	}
	k.store.Store(map[K]V{})
	return k
}
