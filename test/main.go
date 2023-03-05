package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	blanklist []int
	n         int
}

func main() {
	s := Constructor(3, []int{0})
	s.Pick()
	fmt.Println(s)
}

func Constructor(n int, blacklist []int) Solution {
	m := map[int]int{}
	for k, v := range blacklist {
		m[v] = k
	}

	blanklist := make([]int, 2*10000)
	count := 0
	for i := 0; i < n; i++ {
		if _, ok := m[i]; !ok {
			blanklist[count] = i
			count++
		}
	}
	return Solution{
		blanklist: blanklist,
		n:         n,
	}
}

func (this *Solution) Pick() int {
	randIndex := rand.Intn(this.n) % len(this.blanklist)
	return this.blanklist[randIndex]
}
