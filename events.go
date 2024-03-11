package gox

import "sync"

// EventManager is a simple thread-safe dispatcher for managing callbacks in the
// case of events.
//
// Any listener that returns true when called will remove itself. This is good for
// one-shot event listeners.
type EventManager[Payload any] struct {
	lock      sync.RWMutex
	listeners map[uint]func(payload Payload) bool
	lastID    uint
}

// Add adds a new listener to the event manager with the given callback. It
// returns the ID for the listener so that it can be removed later with [Remove]
func (e *EventManager[Payload]) Add(cb func(payload Payload) bool) uint {
	e.lock.RLock()
	defer e.lock.RUnlock()

	e.lastID++
	e.listeners[e.lastID] = cb

	return e.lastID
}

// Remove removes the listener with the given ID
func (e *EventManager[Payload]) Remove(id uint) {
	e.lock.RLock()
	defer e.lock.RUnlock()

	delete(e.listeners, id)
}

// Dispatch sends the payload to all the listeners in the list. If any of the
// listeners return true, then they will be removed from the manager.
func (e *EventManager[Payload]) Dispatch(payload Payload) {
	e.lock.Lock()
	defer e.lock.Unlock()

	for k, cb := range e.listeners {
		if cb(payload) {
			// Returning true means to remove itself
			delete(e.listeners, k)
		}
	}
}

func NewEventManager[Payload any]() EventManager[Payload] {
	return EventManager[Payload]{
		listeners: make(map[uint]func(payload Payload) bool),
		lastID:    1,
	}
}
