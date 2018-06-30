package tests

import (
	"go-collections/linear"
	"testing"
)

var queue = new(linear.Queue)

func TestEnqueue(t *testing.T)  {
	if !queue.IsEmpty() {
		t.Error("expected queue to be empty")
	}

	queue.Enqueue("goat")
	queue.Enqueue("sheep")
	queue.Enqueue("whale")
	queue.Enqueue("pigeon")
	queue.Enqueue("duck")

	if queue.Size() != 5 {
		t.Error("expected size to be 5")
	}
}

func TestQueueFirst(t *testing.T) {
	if queue.First() != "goat" {
		t.Error("expected first to be goat")
	}
}

func TestDequeue(t *testing.T)  {
	if queue.Dequeue() != "goat" {
		t.Error("Expected goat")
	}

	if queue.First() != "sheep" {
		t.Error("expected first to be sheep")
	}
}