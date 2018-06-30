package tests

import (
	"go-collections/linear"
	"testing"
)

var stack  = new(linear.Stack)

func TestPush(t *testing.T)  {
	if !stack.IsEmpty() {
		t.Error("expected stack to be empty")
	}

	stack.Push("goat")
	stack.Push("sheep")
	stack.Push("whale")
	stack.Push("pigeon")
	stack.Push("duck")

	if stack.Size() != 5 {
		t.Error("expected size to be 5")
	}
}

func TestTop(t *testing.T) {
	if stack.Top() != "duck" {
		t.Error("expected first to be goat")
	}
}

func TestPop(t *testing.T)  {
	if stack.Pop() != "duck" {
		t.Error("Expected goat")
	}

	if stack.Top() != "pigeon" {
		t.Error("expected first to be pigeon")
	}
}
