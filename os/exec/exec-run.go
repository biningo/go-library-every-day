package main

import (
	"bytes"
	"log"
	"os/exec"
)

/**
*@Author bingo
*@Date 2020/8/4
*
**/
//指定shell命令
func RunDemo1() {
	cmd := exec.Command("ls", "-l") //在哪个目录下执行就展示哪个目录
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil { //Run会阻塞
		log.Println(err)
	}

	log.Println(stdout.String())
	log.Println(stderr.String())
}

func main() {
	RunDemo1()
}
