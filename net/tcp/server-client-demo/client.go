package main

import (
	"io"
	"log"
	"net"
	"time"
)

/**
*@Author icepan
*@Date 2020/8/12 下午5:24
*@Describe
**/

func main() {
	conn,_:=net.Dial("tcp","localhost:10000")
	log.Println("服务器IP",conn.RemoteAddr())


	for{
		conn.Write([]byte("你好,我是客户端"))
		log.Println("等待服务器回复.......")
		buf:=make([]byte,1024)
		if _,err:=conn.Read(buf);err!=nil{
			if err==io.EOF{
				log.Println("服务器关闭了")
				return
			}
		}
		log.Println("收到：",string(buf))
		time.Sleep(5*time.Second)
	}




}