package geecache

import (
	"fmt"
	"log"
	"sync"
)

// 一个缓存的命名空间，每个 Group 拥有一个唯一的名称
type Group struct {
	name      string
	// 即缓存未命中时获取源数据的回调
	getter    Getter
	// 并发缓存
	mainCache cache
}



// A Getter loads data for a key.
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}



var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// 实例化 Group，并且将 group 存储在全局变量 groups 中
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// GetGroup returns the named group previously created with NewGroup, or
// nil if there's no such group.
func GetGroup(name string) *Group {
	// 使用了只读锁，因为不涉及任何冲突变量的写操作
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}





// Get value for a key from cache
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	// 缓存中存在
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}
	// 缓存中不存在
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

// 分布式场景下从其他节点获取
func (g *Group) getLocally(key string) (ByteView, error) {
	fmt.Println("[GeeCache] getter",key)
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err

	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

// 获取后插入数据
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
