package main

import (
	"fmt"
	"sync"
)

type CuncurrentMap struct {
	mu sync.Mutex

	cache map[int]int
}

func NewCuncurrentMap() *CuncurrentMap {
	return &CuncurrentMap{
		cache: make(map[int]int, 0),
	}
}

func (m *CuncurrentMap) Save(key, val int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache[key] = val
}

func main() {
	cuncurrentMap := NewCuncurrentMap()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			cuncurrentMap.Save(i, i)
		}(i)

	}

	wg.Wait()

	fmt.Println(cuncurrentMap.cache)

}
