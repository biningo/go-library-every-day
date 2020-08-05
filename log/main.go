package main

import (
	"log"
	"os"
)

/**
*@Author icepan
*@Date 2020/8/5
*
**/

//1. New 自定义日志
func LoggerNewDemo(){
	log.Println("输出到标准输出") //默认是输出到标准输出
	f,_:=os.OpenFile("log.txt",os.O_WRONLY|os.O_CREATE,0665)
	defer f.Close()

	logger:=log.New(f,"【MyPrefix】",log.Ltime)
	//以下都可以中途改变
	//logger.SetPrefix()
	//logger.SetFlags()
	logger.Println("天气：晴，心情：好") //写入日志到文件
	//logger.Writer().Write([]byte("今天，多云，心情：好")) 可以直接操作Writer
}

func main() {
	//LoggerNewDemo()
}