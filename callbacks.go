package gameliftgo

import "sync"
import "C"

var mu sync.Mutex
var index int
var fns = make(map[int]interface{})

func register(fn interface{}) int {
	mu.Lock()
	defer mu.Unlock()
	index++
	for fns[index] != nil {
		index++
	}
	fns[index] = fn
	return index
}

func lookup(i int) interface{} {
	mu.Lock()
	defer mu.Unlock()
	return fns[i]
}

func unregister(i int) {
	mu.Lock()
	defer mu.Unlock()
	delete(fns, i)
}

func unregisterAll() {
	for k := range fns {
		unregister(k)
	}
}
