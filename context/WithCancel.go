package main

import (
	"context"
	"log"
	"time"
)

/**
*@Author icepan
*@Date 2020/8/6
*
**/



func Run(ctx context.Context, name string) {

	for {
		select {
		case <-ctx.Done():
			log.Println(name,"运行结束")
			return
		default:
			log.Println(name,"运行中....")
			time.Sleep(time.Millisecond*500)
		}
	}

}

func main() {
	ctx,cancel:=context.WithCancel(context.Background()) //Background 就是返回当前的context
	go Run(ctx,"A")
	go Run(ctx,"B")
	go Run(ctx,"C")
	time.Sleep(time.Second*10)
	cancel() //调用cancel回调函数关闭Done通道，表示其他go程可以结束了
	time.Sleep(time.Hour)

}