package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

/**
*@Author icepan
*@Date 2020/8/9
*
**/

func main() {

	log.Println("服务器启动.....")
	l, e := net.Listen("tcp", "localhost:10000") //监听端口
	if e != nil {
		log.Println(e)
		return
	}
	defer l.Close()

	//不断的等待连接 死循环
	for {
		conn, e := l.Accept() //TCP协议 Accept状态 接受链接
		if e != nil {
			log.Println(e)
			return
		}

		log.Println("连接远程IP：", conn.RemoteAddr())

		go handler(conn) //处理链接
	}

}

func handler(conn net.Conn) {
	defer conn.Close()
	log.Println("客户端IP：",conn.RemoteAddr())
	//不断的处理数据
	for {
		fmt.Println("服务器等待客户端发送数据......")
		buf := make([]byte, 1024)

		// 如果没有可读数据，也就是对方没有发送数据，但是还保持连接，则阻塞
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("客户端已经关闭连接", err) //err是io.EOF
			return
		} else {
			fmt.Println("收到：", string(buf[0:n]))

			// 同理，不可写则阻塞
			_, err := conn.Write([]byte("你好，我是服务器"))
			if err != nil {
				fmt.Println("出现不可知异常", err)
			}
		}
	}
}
