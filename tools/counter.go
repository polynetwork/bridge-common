package tools

import (
	"fmt"
	"sync"
	"time"
)

// Simple timed counter with isolated time windows
type TimedCounter struct {
	interval time.Duration
	value    int
	current  int
	sync.Mutex
}

func NewTimedCounter(interval time.Duration) (*TimedCounter, error) {
	if interval <= 0 {
		return nil, fmt.Errorf("Invalid timed counter duration %v", interval)
	}
	c := &TimedCounter{interval: interval}
	go c.run()
	return c, nil
}

func (c *TimedCounter) Tick(count int) {
	if count == 0 {
		return
	}

	c.Lock()
	c.value += count
	c.Unlock()
}

func (c *TimedCounter) run() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.Lock()
		c.current = c.value
		c.value = 0
		c.Unlock()
	}
}

func (c *TimedCounter) Value() int {
	var v int
	c.Lock()
	v = c.current
	c.Unlock()
	return v
}

func (c *TimedCounter) Tps() float32 {
	v := c.Value()
	duration := c.interval / time.Second
	if duration == 0 {
		return 0
	}
	return float32(v) / float32(duration)
}

// Simple block counter
type BlockCounter struct {
	state             [][2]uint64 // buf [height, time]
	count, head, tail int
	last              uint64
	full              bool
	sync.Mutex
}

func NewBlockCounter(count int) (*BlockCounter, error) {
	if count < 2 {
		return nil, fmt.Errorf("Count too small")
	}
	return &BlockCounter{
		state: make([][2]uint64, count),
		count: count,
		head:  1,
	}, nil
}

func (c *BlockCounter) Tick(height uint64) {
	c.Lock()
	c.Unlock()
	if height == c.last {
		return
	}
	c.last = height
	if !c.full && c.tail == c.count-1 {
		c.full = true
	}
	c.tail = (c.tail + 1) % c.count
	c.state[c.tail][0] = height
	c.state[c.tail][1] = uint64(NowMS())
	if c.full {
		c.head = (c.tail + 1) % c.count
	}
}

func (c *BlockCounter) BlockTime() float32 {
	c.Lock()
	defer c.Unlock()
	h := c.state[c.head%c.count]
	t := c.state[c.tail%c.count]
	duration := float32(t[1] - h[1])
	blocks := t[0] - h[0]
	if blocks == 0 {
		return 0
	}
	return duration / float32(blocks)
}
