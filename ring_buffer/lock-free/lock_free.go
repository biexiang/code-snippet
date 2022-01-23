package lock_free

import (
	"sync/atomic"
	"unsafe"
)

type RingBuffer struct {
	queue      []interface{}
	head, tail uint64
	cap, mask  uint64
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
		queue: make([]interface{}, capacity),
		head:  uint64(0),
		tail:  uint64(0),
		cap:   capacity,
		mask:  capacity - 1,
	}
}

func (c *RingBuffer) EnQueue(value interface{}) bool {
	// EnQueue only 非nil的值
	if value == nil {
		return false
	}

	oldHead := atomic.LoadUint64(&c.head)
	oldTail := atomic.LoadUint64(&c.tail)
	if IsFull(oldHead, oldTail, c.cap) {
		return false
	}

	newTail := (oldTail + 1) & c.mask
	// 判断newTail是否为nil
	// 是否可以直接取unsafe.pointer的地址 ？
	if newTailData := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&c.queue[newTail]))); newTailData != nil {
		return false
	}

	if !atomic.CompareAndSwapUint64(&c.tail, oldTail, newTail) {
		return false
	}
	// https://www.flysnow.org/2017/07/06/go-in-action-unsafe-pointer.html
	// https://segmentfault.com/a/1190000017389782
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&c.queue[newTail])), unsafe.Pointer(&value))
	return true
}

func (c *RingBuffer) DeQueue() (value interface{}, success bool) {
	oldHead := atomic.LoadUint64(&c.head)
	oldTail := atomic.LoadUint64(&c.tail)
	if IsEmpty(oldHead, oldTail) {
		return nil, false
	}

	newHead := (oldHead + 1) & c.mask
	headData := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&c.queue[newHead])))
	if headData == nil {
		return nil, false
	}

	if !atomic.CompareAndSwapUint64(&c.head, oldHead, newHead) {
		return nil, false
	}

	// 原数据置为nil
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&c.queue[newHead])), nil)
	return *(*interface{})(headData), true
}

func IsEmpty(head, tail uint64) bool {
	return head == tail
}

func IsFull(head, tail, cap uint64) bool {
	return tail-head == cap-1 || head-tail == 1
}
