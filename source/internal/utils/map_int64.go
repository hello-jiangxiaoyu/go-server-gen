package utils

import "sync"

type MapInt64[T any] struct {
	m sync.Map
}

// Get 获取key对应的值
func (w *MapInt64[T]) Get(key int64) *T {
	v, ok := w.m.Load(key)
	if !ok {
		return nil
	}
	res, ok := v.(*T)
	if !ok {
		return nil
	}

	return res
}

// Set 插入key
func (w *MapInt64[T]) Set(key int64, value *T) {
	w.m.Store(key, value)
}

// Delete 删除key
func (w *MapInt64[T]) Delete(key int64) {
	w.m.Delete(key)
}

// DeleteAll 清空
func (w *MapInt64[T]) DeleteAll() {
	w.m.Range(func(key, _ any) bool {
		w.m.Delete(key)
		return true
	})
}

func (w *MapInt64[T]) Range(fn func(int64, *T) bool) {
	w.m.Range(func(k, v any) bool {
		key, ok := k.(int64)
		if !ok {
			return false
		}
		value, ok := v.(*T)
		if !ok {
			return false
		}
		if !fn(key, value) {
			return false
		}
		return true
	})
}
