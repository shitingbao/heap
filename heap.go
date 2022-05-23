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
type heap struct {
	List   []int
	Symbol bool // true 代表
}

// NewMaxHeap 可传入多个切片构造一个小根堆
func NewMaxHeap(list []int) (*MaxHeap, error) {
	h := &heap{Symbol: true, List: make([]int, len(list))}
	h.construct(list)
	return &MaxHeap{h}, nil
}

// NewMinHeap 可传入多个切片构造一个小根堆
func NewMinHeap(list []int) (*MinHeap, error) {
	h := &heap{Symbol: false, List: make([]int, len(list))}
	h.construct(list)
	return &MinHeap{h}, nil
}

// construct 堆构造函数
func (h *heap) construct(list []int) {
	for i, v := range list {
		h.List[i] = v
		h.upSort(i)
	}
}

func (m *heap) GetValue() (int, error) {
	return m.getValue()
}

func (m *heap) getValue() (int, error) {
	if len(m.List) == 0 {
		return 0, errors.New("heap is no val")
	}
	p := m.List[0]
	m.List[0], m.List[len(m.List)-1] = m.List[len(m.List)-1], m.List[0]
	m.List = m.List[:len(m.List)-1]
	m.lowSort(0)
	return p, nil
}

func (m *heap) PutValue(val int) { m.putValue(val) }

func (m *heap) putValue(val int) {
	m.List = append(m.List, val)
	m.upSort(len(m.List) - 1)

}

// 上浮排序,symbol true 为小根队
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
		case m.Symbol:
			if (m.List)[index] < (m.List)[flag] {
				(m.List)[index], (m.List)[flag] = (m.List)[flag], (m.List)[index]
			}
		default:
			if (m.List)[index] > (m.List)[flag] {
				(m.List)[index], (m.List)[flag] = (m.List)[flag], (m.List)[index]
			}
		}
		flag = index
	}
}

// 下浮,symbol true 为小根队
func (m *heap) lowSort(flag int) {
	for {
		left := flag*2 + 1        // 获取左孩子
		if left > len(m.List)-1 { // 无左孩子说明结束
			break
		}
		right := flag*2 + 2 // 获取右孩子
		index := left
		switch {
		case m.Symbol:
			// 有右孩子的情况下，判断左右孩子的大小
			if right <= len(m.List)-1 && (m.List)[left] < (m.List)[right] {
				index = right
			}
			if (m.List)[index] > (m.List)[flag] {
				(m.List)[index], (m.List)[flag] = (m.List)[flag], (m.List)[index]
			}
		default:
			// 有右孩子的情况下，判断左右孩子的大小
			if right <= len(m.List)-1 && (m.List)[left] > (m.List)[right] {
				index = right
			}
			if (m.List)[index] < (m.List)[flag] {
				(m.List)[index], (m.List)[flag] = (m.List)[flag], (m.List)[index]
			}
		}
		flag = index
	}
}
