package main

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"
)

/**
*@Author icepan
*@Date 2020/8/6
*
**/
var wg sync.WaitGroup
func main() {
	//context.WithDeadline() 指定一个时间点后关闭通道
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5) //5s之后关闭Done通道
	for i:=0;i<10;i++{
		go RunInt(ctx,i)
	}
	_=<-ctx.Done()
	log.Println("运行结束")
	wg.Wait()
}

func RunInt(c context.Context,i int){
	wg.Add(1)
	for {
		select {
		case <-c.Done():
			log.Println(strconv.Itoa(i)+":结束")
			wg.Done()
			return
		default:
			log.Println(strconv.Itoa(i)+":运行中...")
			time.Sleep(time.Millisecond * 500)
		}
	}
}
