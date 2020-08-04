package main

import (
	"log"
	"sync"
	"sync/atomic"
)

/**
*@Author bingo
*@Date 2020/8/4
*
**/

//举例int32类型 其他类型一致

//一、 Load Store  分别对应get和set操作
func LoadStoreDemo() {
	var val int32
	atomic.StoreInt32(&val, 10) //无返回值
	j := atomic.LoadInt32(&val) //10
	log.Println(j)
}

//二、 Add  增加  减少的话传入负数即可
func AddDemo() {
	wg1 := sync.WaitGroup{}
	var val int32

	//1.并发安全
	for i := 0; i < 100; i++ {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			atomic.AddInt32(&val, 10)
		}()
	}
	wg1.Wait()
	log.Println(val) //1000

	//2.并发不安全
	val = 0
	wg2 := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			val += 10
		}()
	}
	wg2.Wait()
	log.Println(val) // 可能出现小于1000的情况
}

//三、交换 CompareAndSwap
func CompareSwapDemo() {

	//可以把val当做开关 并发执行时  防止重复关闭或者打开
	var val int32 = 0
	const CLOSE int32 = 0
	const OPEN int32 = 1

	atomic.CompareAndSwapInt32(&val, CLOSE, OPEN) // val==old==0 所以交换 原来是关的 现在打开
	log.Println(val)                              //OPEN 1

	atomic.CompareAndSwapInt32(&val, OPEN, CLOSE) // val==old 原来是开的 现在关闭
	log.Println(val)                              //CLOSE 0

	atomic.CompareAndSwapInt32(&val, OPEN, CLOSE) //val!=OPEN 所以不交换
	log.Println(val)                              //CLOSE 0
}

func main() {
	LoadStoreDemo()
	AddDemo()
	CompareSwapDemo()
}
