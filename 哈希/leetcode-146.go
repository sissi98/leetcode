package 哈希

type Cache struct {
	key   int
	value int
}
type LRUCache struct {
	capacity int //缓存容量
	cache    []Cache
	mp       map[int]int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, []Cache{}, map[int]int{}}
}

func (this *LRUCache) Get(key int) int {
	if m, ok := this.mp[key]; ok {
		for k, v := range this.cache {
			if v.key == key {
				this.cache = append(this.cache[:k], this.cache[k+1:]...)
				this.cache = append(this.cache, Cache{key, v.value})
				break
			}
		}
		return m
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// 缓存中没有当前值
	if _, ok := this.mp[key]; !ok {
		if len(this.cache) == this.capacity {
			delete(this.mp, this.cache[0].key)
			this.cache = this.cache[1:]
		}
	} else {
		//缓存中有当前值
		for k, v := range this.cache {
			if v.key == key {
				this.cache = append(this.cache[:k], this.cache[k+1:]...)
				break
			}
		}
	}
	this.mp[key] = value
	this.cache = append(this.cache, Cache{key, value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
