package _46_LRUCache

import (
	"testing"
)

//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4

func TestLRU(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	v1 := lru.Get(1)
	assertEqual(t, v1, 1)
	lru.Put(3, 3)
	v2 := lru.Get(2)
	assertEqual(t, v2, -1)
	lru.Put(4, 4)
	v3 := lru.Get(1)
	assertEqual(t, v3, -1)
	v4 := lru.Get(3)
	assertEqual(t, v4, 3)
	v5 := lru.Get(4)
	assertEqual(t, v5, 4)
}

func assertEqual(t *testing.T, v1 int, i int) {
	if v1 != i {
		t.Errorf("assertEqual %v != %v", v1, i)
	}
}
