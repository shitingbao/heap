package heap

import (
	"log"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	list := []int{9, 3, 7, 6, 5, 1, 10, 2}
	h, err := NewMaxHeap(list)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("first:", h.List)
	log.Println(h.getValue())
	log.Println("h:", h.List)
	h.putValue(8)
	h.putValue(14)
	log.Println(h.getValue())
	log.Println("h:", h.List)
}

func TestMinHeap(t *testing.T) {
	list := []int{9, 3, 7, 6, 5, 1, 10, 2}
	h, err := NewMinHeap(list)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("first:", h.List)
	log.Println(h.getValue())
	log.Println("h:", h.List)
	h.putValue(8)
	h.putValue(14)
	log.Println(h.getValue())
	log.Println("h:", h.List)
}
