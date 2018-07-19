package tests

import (
	"testing"
	"go-collections/counter"
	"sync"
)

var count = counter.GetInstance()

func TestGetInstance(t *testing.T)  {
	if count == nil {
		t.Error("counter should not be nil")
	}
}

func TestCounter_Increment(t *testing.T) {
	var inc sync.WaitGroup

	for i := 0; i < 5; i++ {

		inc.Add(1)
		go func() {
			defer inc.Done()
			count.Increment()
		}()
	}

	inc.Wait()

	if count.GetValue() != 5 {
		t.Error("value should be 5")
	}

}

func TestCounter_Decrement(t *testing.T)  {
	var dec sync.WaitGroup
	for i := 0; i < 3; i++  {

		dec.Add(1)
		go func() {
			defer dec.Done()
			count.Decrement()
		}()
	}

	dec.Wait()

	if count.GetValue() != 2 {
		t.Error("value should be 2")
	}
}

func TestGetValue(t *testing.T)  {
	if count.GetValue() != 2 {
		t.Error("value should be 2")
	}
}