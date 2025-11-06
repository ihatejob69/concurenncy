package counter

import "sync"

type Counter struct {
	mu sync.Mutex
	v  int
}

func (c *Counter) Inc() {
	// TODO: реализовать инкремент с использованием мьютекса
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	// TODO: вернуть значение с учётом блокировки
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}
