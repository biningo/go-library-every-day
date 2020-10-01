package main

import (
	"log"
	"sync"
	"time"
)

/**
*@Author icepan
*@Date 2020/10/1 上午10:12
*@Describe
**/

func T1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			log.Println(v)
		}(i)
	}
	wg.Wait()
}

func T2() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		log.Println(i)
		time.Sleep(time.Second)
		wg.Done()
	}
	wg.Wait()
}

func main() {
	//T1()
	T2()
}
