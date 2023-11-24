package heap

import (
	"errors"
	"sync"
)

// MinHeap 小根堆
type MinHeap struct {
	*heap
}

// MaxHeap 大根堆
type MaxHeap struct {
	*heap
}

// 基本的队列结构
// 用切片构造基本数据类型，索引对应树内节点序号
type heap struct {
	list   []int
	symbol bool // true 代表大根堆
	lock   sync.Locker
}

// NewMaxHeap 可传入切片构造一个大根堆
func NewMaxHeap(list []int) (*MaxHeap, error) {
	h := &heap{symbol: true, list: make([]int, len(list)), lock: NewSpinLock()}
	h.construct(list)
	return &MaxHeap{h}, nil
}

// NewMinHeap 可传入切片构造一个小根堆
func NewMinHeap(l []int) (*MinHeap, error) {
	h := &heap{symbol: false, list: make([]int, len(l)), lock: NewSpinLock()}
	h.construct(l)
	return &MinHeap{h}, nil
}

// construct 堆构造函数
// 将元素依次加入堆中构造完整的堆
func (h *heap) construct(list []int) {
	for i, v := range list {
		h.list[i] = v
		h.upSort(i)
	}
}

func (h *heap) TryGetValue() (int, error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	return h.tryGetValue()
}

// 反馈堆顶的值，并下浮，重新形成新的堆
func (h *heap) tryGetValue() (int, error) {
	if len(h.list) == 0 {
		return 0, errors.New("heap is no val")
	}
	p := h.list[0]
	return p, nil
}

func (h *heap) GetValue() (int, error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	return h.getValue()
}

// 反馈堆顶的值，并下浮，重新形成新的堆
func (h *heap) getValue() (int, error) {
	p, err := h.tryGetValue()
	if err != nil {
		return 0, err
	}
	h.list[0], h.list[len(h.list)-1] = h.list[len(h.list)-1], h.list[0]
	h.list = h.list[:len(h.list)-1]
	h.lowSort(0)
	return p, nil
}

func (h *heap) PutValue(val int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.putValue(val)
}

// 加入一个值，并上浮形成新的堆
func (h *heap) putValue(val int) {
	h.list = append(h.list, val)
	h.upSort(len(h.list) - 1)
}

func (h *heap) List() []int {
	return h.list
}

// 上浮排序,symbol true 为大根队
// 获取父节点序号，比较大小并交换
func (h *heap) upSort(flag int) {
	for {
		if flag <= 0 {
			break
		}
		index := (flag - 1) / 2
		if (flag)%2 == 0 {
			index = (flag - 2) / 2
		}
		switch {
		case h.symbol:
			if (h.list)[index] < (h.list)[flag] {
				(h.list)[index], (h.list)[flag] = (h.list)[flag], (h.list)[index]
			}
		default:
			if (h.list)[index] > (h.list)[flag] {
				(h.list)[index], (h.list)[flag] = (h.list)[flag], (h.list)[index]
			}
		}
		flag = index
	}
}

// 下浮,symbol true 为大根队
// 获取子节点，判断大小并交换，注意左右节点是否存在
func (h *heap) lowSort(flag int) {
	for {
		left := flag*2 + 1        // 获取左孩子
		if left > len(h.list)-1 { // 无左孩子说明结束
			break
		}
		right := flag*2 + 2 // 获取右孩子
		index := left
		switch {
		case h.symbol:
			// 有右孩子的情况下，判断左右孩子的大小
			if right <= len(h.list)-1 && (h.list)[left] < (h.list)[right] {
				index = right
			}
			if (h.list)[index] > (h.list)[flag] {
				(h.list)[index], (h.list)[flag] = (h.list)[flag], (h.list)[index]
			}
		default:
			// 有右孩子的情况下，判断左右孩子的大小
			if right <= len(h.list)-1 && (h.list)[left] > (h.list)[right] {
				index = right
			}
			if (h.list)[index] < (h.list)[flag] {
				(h.list)[index], (h.list)[flag] = (h.list)[flag], (h.list)[index]
			}
		}
		flag = index
	}
}
