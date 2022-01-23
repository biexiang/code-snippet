package no_lock

type RingBuffer struct {
	Queue      []interface{}
	Head, Tail uint64
	Cap, Mask  uint64
}

func findPowerOfTwo(givenMum uint64) uint64 {
	givenMum--
	givenMum |= givenMum >> 1
	givenMum |= givenMum >> 2
	givenMum |= givenMum >> 4
	givenMum |= givenMum >> 8
	givenMum |= givenMum >> 16
	givenMum |= givenMum >> 32
	givenMum++
	return givenMum
}

func Constructor(k int) RingBuffer {
	capacity := findPowerOfTwo(uint64(k))
	return RingBuffer{
		Queue: make([]interface{}, capacity),
		Head:  uint64(0),
		Tail:  uint64(0),
		Cap:   capacity,
		Mask:  capacity - 1,
	}
}

func (c *RingBuffer) EnQueue(value interface{}) bool {
	if c.IsFull() {
		return false
	}
	newTail := (c.Tail + 1) & c.Mask
	c.Tail = newTail
	c.Queue[newTail] = value
	return true
}

func (c *RingBuffer) DeQueue() (value interface{}, success bool) {
	if c.IsEmpty() {
		return nil, false
	}
	newHead := (c.Head + 1) & c.Mask
	c.Head = newHead
	return c.Queue[newHead], true
}

func (c *RingBuffer) IsEmpty() bool {
	return c.Head == c.Tail
}

func (c *RingBuffer) IsFull() bool {
	return c.Tail-c.Head == c.Cap-1 || c.Head-c.Tail == 1
}

//func (c *RingBuffer) Front() interface{} {
//	if c.IsEmpty() {
//		return nil
//	}
//	return c.queue[c.head]
//}
//
//func (c *RingBuffer) Rear() interface{} {
//	if c.IsEmpty() {
//		return nil
//	}
//	return c.queue[c.tail]
//}
