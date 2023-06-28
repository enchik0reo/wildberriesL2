package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(chs ...<-chan int) <-chan int {
	c := make(chan int)

	go func() {
		count := len(chs)
		mp := make(map[int]struct{})
		for {
			if count == 0 {
				close(c)
				break
			}
			for i, ch := range chs {
				select {
				case v, ok := <-ch:
					if !ok {
						if _, ok := mp[i]; !ok {
							count--
							mp[i] = struct{}{}
						}
					} else {
						c <- v
					}
				default:
					continue
				}
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 4, 7, 10)
	b := asChan(2, 5, 8, 11)
	d := asChan(3, 6, 9, 12)
	c := merge(a, b, d)
	for v := range c {
		fmt.Println(v)
	}
}
