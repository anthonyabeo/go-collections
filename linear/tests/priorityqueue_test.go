package tests

import (
	"go-collections/linear"
	"testing"
)

var pq = linear.NewPQ()

func TestInsertAndMin(t *testing.T) {

	pq.Insert(2, "Iron Man")
	pq.Insert(1, "Captain America")
	pq.Insert(5, "Hawk Eye")
	pq.Insert(4, "Black Widow")
	pq.Insert(3, "Hulk")

	if pq.Size() != 5 {
		t.Error("number of avengers should be 5")
	}

	if pq.Min().Value() != "Captain America" {
		t.Error("Min should be Cap")
	}

	//fmt.Println(pq.Heap()[0])
}

func TestRemoveMin(t *testing.T) {
	entry := pq.RemoveMin()
	if pq.Size() != 4 {
		t.Error("number of avengers should be 4")
	}

	if entry.Value() != "Captain America" {
		t.Error("removed min should be Cap")
	}
	//fmt.Println(pq.Heap()[3])
}
