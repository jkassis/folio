// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

type LLM struct {
	key  string
	prev *LLM
	next *LLM
	val  string
}

type LRU struct {
	keyCountMax int
	cache       map[string]*LLM
	qStart      *LLM
	qEnd        *LLM
}

func MakeLRU(keyCountMax int) *LRU {
	return &LRU{
		keyCountMax: keyCountMax,
		cache:       make(map[string]*LLM, 0),
	}
}

func (lru *LRU) Put(k, v string) {
	var llm *LLM
	var found bool
	llm, found = lru.cache[k]
	if !found {
		llm = &LLM{
			key:  k,
			prev: lru.qEnd,
			val:  v,
		}

		if len(lru.cache) >= lru.keyCountMax {
			delete(lru.cache, lru.qStart.key)
			lru.qStart = lru.qStart.next
			lru.qStart.prev = nil
		}
		lru.cache[k] = llm

		if lru.qEnd != nil {
			lru.qEnd.next = llm
			llm.prev = lru.qEnd
		}
		lru.qEnd = llm

		if lru.qStart == nil {
			lru.qStart = llm
		}
	} else {
		llm.prev.next = llm.next
		llm.next.prev = llm.prev
		llm.next = nil
		lru.qEnd.next = llm
		llm.prev = lru.qEnd
		lru.qEnd = llm

		llm.val = v
	}
}

// Existing:  (k1, v1), (k2, v2), (k3, v3), (k4, v4)
// Put (k2, newv2)

func (lru *LRU) Get(k, v string) {
	// copy-paste from Put
}

func main() {
	l := MakeLRU(4)
	l.Put("k1", "v1")
	l.Put("k2", "v2")
	l.Put("k3", "v3")
	l.Put("k4", "v4")
	l.Put("k2", "v2new")

	for k, v := range l.cache {
		fmt.Printf("%s : %s\n", k, v.val)
	}

	fmt.Printf("\n")
	n := l.qStart
	for {
		fmt.Printf("%s : %s\n", n.key, n.val)
		n = n.next
		if n == nil {
			break
		}
	}
}
