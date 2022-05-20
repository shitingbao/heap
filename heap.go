package main

import (
	"errors"
)

type MinHeap []int

func NewMinHeap(lists ...[]int) (*MinHeap, error) {
	resList := &MinHeap{}
	if len(lists) > 0 {
		for _, list := range lists {
			resList = construct(true, list)
		}
	}
	return resList, nil
}

// construct 堆构造函数
func construct(symbol bool, list []int) *MinHeap {
	hp := make(MinHeap, len(list))
	for i, v := range list {
		hp[i] = v
		hp.upSort(i, symbol)
	}
	return &hp
}

func (m *MinHeap) getValue() (int, error) {
	if len(*m) == 0 {
		return 0, errors.New("heap is no val")
	}
	p := (*m)[0]
	(*m)[0], (*m)[len(*m)-1] = (*m)[len(*m)-1], (*m)[0]
	*m = (*m)[:len(*m)-1]
	m.lowSort(0, true)
	return p, nil
}

func (m *MinHeap) putValue(val int) {
	*m = append(*m, val)
	m.upSort(len(*m)-1, true)

}

// 上浮排序,symbol true 为小根队
func (m *MinHeap) upSort(flag int, symbol bool) {
	for {
		if flag <= 0 {
			break
		}
		index := (flag - 1) / 2
		if (flag)%2 == 0 {
			index = (flag - 2) / 2
		}
		switch {
		case symbol:
			if (*m)[index] > (*m)[flag] {
				(*m)[index], (*m)[flag] = (*m)[flag], (*m)[index]
			}
		default:
			if (*m)[index] < (*m)[flag] {
				(*m)[index], (*m)[flag] = (*m)[flag], (*m)[index]
			}
		}

		flag = index
	}
}

// 下浮,symbol true 为小根队
func (m *MinHeap) lowSort(flag int, symbol bool) {
	for {
		left := flag*2 + 1    // 获取左孩子
		if left > len(*m)-1 { // 无左孩子说明结束
			break
		}
		right := flag*2 + 2 // 获取右孩子

		index := left
		// 有右孩子的情况下，判断左右孩子的大小
		if right <= len(*m)-1 && (*m)[left] > (*m)[right] {
			index = right
		}
		switch {
		case symbol:
			if (*m)[index] < (*m)[flag] {
				(*m)[index], (*m)[flag] = (*m)[flag], (*m)[index]
			}
		default:
			if (*m)[index] > (*m)[flag] {
				(*m)[index], (*m)[flag] = (*m)[flag], (*m)[index]
			}
		}

		flag = index
	}
}
