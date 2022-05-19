package heap

import (
	"log"
	"sync"
)

type Heap interface {
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

func NewMinHeap(lists ...[]int) (*MinHeap, error) {
	resList := []int{}
	if len(lists) > 0 {
		for _, list := range lists {
			resList = construct(list)
		}
	}
	log.Println(resList)
	return &MinHeap{
		list: []heapElement{},
		lock: &sync.Mutex{},
	}, nil
}

func (h *MinHeap) GetValue() {

}

func (h *MinHeap) PutValue(obj, flag interface{}) {

}

// construct 堆构造函数
func construct(list []int) []int {
	resList := make([]int, len(list))
	for i, v := range list {
		resList[i] = v
		sort(i, resList)
	}
	return resList
}

// 上浮排序（小根）
func sort(flag int, list []int) {
	for {
		if flag <= 0 {
			break
		}
		index := (flag - 1) / 2
		if (flag)%2 == 0 {
			index = (flag - 2) / 2
		}
		if list[index] > list[flag] {
			list[index], list[flag] = list[flag], list[index]
		}
		flag = index
	}
}

// 下浮（小根）
func lowSort(flag int, list []int) {
	for {
		left := flag*2 + 1
		if left > len(list)-1 {
			break
		}
		right := flag*2 + 2

		index := left
		if right <= len(list)-1 && list[left] > list[right] {
			index = right
		}
		if list[index] < list[flag] {
			list[index], list[flag] = list[flag], list[index]
		}
		flag = index
	}
}
