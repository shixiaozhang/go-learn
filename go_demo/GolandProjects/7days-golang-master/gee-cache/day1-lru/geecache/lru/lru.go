package lru

import "container/list"

// Cache is a LRU cache. 并发访问是不安全的。
type Cache struct {
	// 允许使用的最大内存
	maxBytes int64
	// 是当前已使用的内存
	nbytes   int64
	ll       *list.List
	// 键是字符串，值是双向链表中对应节点的指针。
	cache    map[string]*list.Element
	// 可选，某条记录被移除时的回调函数，可以为 nil
	OnEvicted func(key string, value Value)
}

// 是双向链表节点的数据类型
type entry struct {
	key   string
	value Value
}

// Value使用Len来计算需要多少字节
type Value interface {
	Len() int
}

// New是缓存的构造函数
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 新增数据
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)				// ???????????????????????????????????????????
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		// 向队尾添加新节点
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// 查找	第一步是从字典中找到对应的双向链表的节点，第二步，将该节点移动到队尾。
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		// 将元素ele移到列表ll的前面（队尾）
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// 删除队首的节点
func (c *Cache) RemoveOldest() {
	// 返回列表中最后的元素
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Len缓存条目数
func (c *Cache) Len() int {
	return c.ll.Len()
}
