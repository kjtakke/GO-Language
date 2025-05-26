package utils

import (
    "errors"
    "sync"
)

// SafeQueue is a thread-safe FIFO queue backed by a slice and protected by a mutex.
type SafeQueue struct {
    mu    sync.Mutex
    items []interface{}
}

// NewSafeQueue creates and returns a new SafeQueue.
func NewSafeQueue() *SafeQueue {
    return &SafeQueue{
        items: make([]interface{}, 0),
    }
}

// Enqueue adds an item to the end of the queue.
func (q *SafeQueue) Enqueue(item interface{}) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
// Returns an error if the queue is empty.
func (q *SafeQueue) Dequeue() (interface{}, error) {
    q.mu.Lock()
    defer q.mu.Unlock()

    if len(q.items) == 0 {
        return nil, errors.New("queue is empty")
    }

    item := q.items[0]
    q.items = q.items[1:]
    return item, nil
}

// Peek returns the first item in the queue without removing it.
func (q *SafeQueue) Peek() (interface{}, error) {
    q.mu.Lock()
    defer q.mu.Unlock()

    if len(q.items) == 0 {
        return nil, errors.New("queue is empty")
    }

    return q.items[0], nil
}

// Length returns the number of items in the queue.
func (q *SafeQueue) Length() int {
    q.mu.Lock()
    defer q.mu.Unlock()
    return len(q.items)
}

// Clear removes all items from the queue.
func (q *SafeQueue) Clear() {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.items = []interface{}{}
}

// Values returns a copy of the queue's items.
func (q *SafeQueue) Values() []interface{} {
    q.mu.Lock()
    defer q.mu.Unlock()
    return append([]interface{}(nil), q.items...)
}
