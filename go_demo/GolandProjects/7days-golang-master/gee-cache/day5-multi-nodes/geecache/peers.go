package geecache

// PeerPicker is the interface that must be implemented to locate
// the peer that owns a specific key.
// 根据传入的 key 选择相应节点
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter is the interface that must be implemented by a peer.
// 从对应 group 查找缓存值 HTTP 客户端
type PeerGetter interface {
	Get(group string, key string) ([]byte, error)
}
