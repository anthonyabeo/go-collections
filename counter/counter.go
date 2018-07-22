package counter

import (
	"sync"
)

var instance *Counter

type Counter struct {
	value int
	lock *sync.Mutex
}

func GetInstance() *Counter {
	if instance == nil {
		instance = &Counter{
			0,
			new(sync.Mutex),
		}
	}

	return instance
}

func (c *Counter) Increment()  {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.value++
	//fmt.Printf("Incrementing: %d\n", c.value)
}

func (c *Counter) Decrement()  {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.value--
	//fmt.Printf("Decrementing: %d\n", c.value)
}

func (c *Counter) GetValue() int {
	c.lock.Lock()
	defer c.lock.Unlock()

	var val = c.value

	return val
}