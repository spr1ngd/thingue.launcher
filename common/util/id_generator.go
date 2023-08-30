package util

import "sync"

type IDGenerator struct {
	counter     int
	invalidIDs  map[int]bool
	invalidLock sync.Mutex
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{
		counter:    1,
		invalidIDs: make(map[int]bool),
	}
}

func (g *IDGenerator) GenerateID() int {
	g.invalidLock.Lock()
	defer g.invalidLock.Unlock()

	// Reuse invalidated IDs if available
	for id := range g.invalidIDs {
		delete(g.invalidIDs, id)
		return id
	}

	// Generate a new ID
	id := g.counter
	g.counter++
	return id
}

func (g *IDGenerator) InvalidateID(id int) {
	g.invalidLock.Lock()
	defer g.invalidLock.Unlock()

	g.invalidIDs[id] = true
}
