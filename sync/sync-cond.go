package main

import (
	"log"
	"sync"
	"time"
)

/**
*@Author bingo
*@Date 2020/8/2
*
**/

//sync.cond 条件变量 实现多个go程之间的同步
func DemoCond1(){
	var cond = sync.NewCond(new(sync.Mutex))

	for i:=0;i<10;i++{
		go func(x int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait() //阻塞等待通知 必须加锁之后才可以调用
			log.Println(x)
		}(i)
	}

	time.Sleep(time.Millisecond*500)
	log.Println("Signal....")
	cond.Signal() //随机唤醒一个

	time.Sleep(time.Millisecond*500)
	log.Println("Signal....")
	cond.Signal() //随机唤醒一个

	time.Sleep(time.Millisecond*500)
	log.Println("Signal....")
	cond.Signal() //随机唤醒一个


	time.Sleep(time.Second)
	log.Println("Broadcast....")
	cond.Broadcast()//全部唤醒 Wait的go程

	time.Sleep(time.Hour)
}

func main() {
	DemoCond1()
}
