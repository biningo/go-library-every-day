package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

/**
*@Author icepan
*@Date 2020/11/21 下午5:16
*@Describe
**/

func main() {
	c := make(chan int)
	sigs := make(chan os.Signal)
	//https://www.jianshu.com/p/f6dfbf51c541
	//SIGINT(ctrl + c )终止前台进程  SIGTREM(kill pid)终止进程,会被阻塞 SIGKILL(kill -9 pid)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //将 SIGINT(ctrl+c) 改变为 SIGTREM(kill pid)
	go func() {
		<-sigs
		log.Println("bye bye....")
		c <- 1
	}()
	<-c
}
