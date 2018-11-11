package main

import (
	"fmt"
	"sync"
)

type cwork struct {
	in chan int
	wg *sync.WaitGroup
}

func main() {

	c := [10]chan int{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c[i] = make(chan int)
		go func(i int, c chan int, wg *sync.WaitGroup) {
			for n := range c {
				fmt.Printf("第 %d 个worker为:%d\n", i, n)
				wg.Done()
			}
		}(i, c[i], &wg)

	}

	for i, worker := range c {
		worker <- i
		wg.Add(1)
	}

	wg.Wait()
}
