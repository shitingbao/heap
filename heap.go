package heap

import (
	"errors"
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
}

// NewMaxHeap 可传入切片构造一个大根堆
func NewMaxHeap(list []int) (*MaxHeap, error) {
	h := &heap{symbol: true, list: make([]int, len(list))}
	h.construct(list)
	return &MaxHeap{h}, nil
}

// NewMinHeap 可传入切片构造一个小根堆
func NewMinHeap(l []int) (*MinHeap, error) {
	h := &heap{symbol: false, list: make([]int, len(l))}
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

func (m *heap) GetValue() (int, error) {
	return m.getValue()
}

// 反馈堆顶的值，并下浮，重新形成新的堆
func (m *heap) getValue() (int, error) {
	if len(m.list) == 0 {
		return 0, errors.New("heap is no val")
	}
	p := m.list[0]
	m.list[0], m.list[len(m.list)-1] = m.list[len(m.list)-1], m.list[0]
	m.list = m.list[:len(m.list)-1]
	m.lowSort(0)
	return p, nil
}

func (m *heap) PutValue(val int) { m.putValue(val) }

// 加入一个值，并上浮形成新的堆
func (m *heap) putValue(val int) {
	m.list = append(m.list, val)
	m.upSort(len(m.list) - 1)
}

func (m *heap) List() []int {
	return m.list
}

// 上浮排序,symbol true 为大根队
// 获取父节点序号，比较大小并交换
func (m *heap) upSort(flag int) {
	for {
		if flag <= 0 {
			break
		}
		index := (flag - 1) / 2
		if (flag)%2 == 0 {
			index = (flag - 2) / 2
		}
		switch {
		case m.symbol:
			if (m.list)[index] < (m.list)[flag] {
				(m.list)[index], (m.list)[flag] = (m.list)[flag], (m.list)[index]
			}
		default:
			if (m.list)[index] > (m.list)[flag] {
				(m.list)[index], (m.list)[flag] = (m.list)[flag], (m.list)[index]
			}
		}
		flag = index
	}
}

// 下浮,symbol true 为大根队
// 获取子节点，判断大小并交换，注意左右节点是否存在
func (m *heap) lowSort(flag int) {
	for {
		left := flag*2 + 1        // 获取左孩子
		if left > len(m.list)-1 { // 无左孩子说明结束
			break
		}
		right := flag*2 + 2 // 获取右孩子
		index := left
		switch {
		case m.symbol:
			// 有右孩子的情况下，判断左右孩子的大小
			if right <= len(m.list)-1 && (m.list)[left] < (m.list)[right] {
				index = right
			}
			if (m.list)[index] > (m.list)[flag] {
				(m.list)[index], (m.list)[flag] = (m.list)[flag], (m.list)[index]
			}
		default:
			// 有右孩子的情况下，判断左右孩子的大小
			if right <= len(m.list)-1 && (m.list)[left] > (m.list)[right] {
				index = right
			}
			if (m.list)[index] < (m.list)[flag] {
				(m.list)[index], (m.list)[flag] = (m.list)[flag], (m.list)[index]
			}
		}
		flag = index
	}
}
