package main

import "log"

var cache = make(map[int64]int64, 100)

// The result of F(100) is 24382819596721629
func main() {
	log.Printf("F(100): %d\n", F(100))
}

// F function
func F(n int64) int64 {
	if n < 3 {
		if _, ok := cache[n]; !ok {
			cache[n] = 1
		}
		return 1
	}
	if result, ok := cache[n]; ok {
		return result
	}
	v1 := F(n - 1)
	v2 := F(n - 3)
	cache[n] = v1 + v2
	return cache[n]
}