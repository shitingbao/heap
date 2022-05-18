package heap

import "sync"

type Heap interface {
	sort()
	GetValue()
	PutValue(obj, flag interface{}) // obj,存储的内容，flag 根据这个排序
}

type heapElement struct {
	Value interface{}
	Flag  interface{}
}

var _ Heap = (*MinHeap)(nil)

type MinHeap struct {
	list []heapElement
	lock *sync.Mutex
}

// 递减
func sortLow(list []heapElement) {}

func (h *MinHeap) sort() {
	h.lock.Lock()
	defer h.lock.Unlock()
	sortLow(h.list)
}

func (h *MinHeap) GetValue() {
	h.list[0], h.list[len(h.list)-1] = h.list[len(h.list)-1], h.list[0]
	h.sort()
}

func (h *MinHeap) PutValue(obj, flag interface{}) {
	h.list = append(h.list, heapElement{
		Value: obj,
		Flag:  flag,
	})
	h.sort()
}

type MaxHeap struct {
	list []heapElement
	lock *sync.Mutex
}

// 递增
func sortUp(list []heapElement) {}

func NewMaxHeap() {

}

func NewMaixHeap() Heap {
	return &MinHeap{
		list: []heapElement{},
		lock: &sync.Mutex{},
	}
}
