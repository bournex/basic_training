package cache

// LFU
type VLfu interface {
	Init(int64)
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	ToString() string
}

type lfuNode struct {
}

// LFU 最近最少用缓存实现思路
// 最近 - 节点记录时间戳，每次访问时更新其时间戳
// 最少用 - 节点记录访问次数
// 缓存策略，
type lfu struct {
	max int64
}
