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
//1.Hello world
func RunDemo1() {
	cmd := exec.Command("ls", "-l") //在哪个目录下执行就展示哪个目录
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	//cmd.Start()和cmd.Wait()的组合
	if err := cmd.Run(); err != nil { //Run会阻塞
		log.Println(err)
	}

	log.Println(stdout.String())
	log.Println(stderr.String())
}

//2.设置命令运行时的目录 默认在当前目录下
func RunDemoDir()  {
	cmd:=exec.Command("ls")
	cmd.Dir="/" //设置cmd运行时的目录 默认就是当前目录
	var msg bytes.Buffer
	cmd.Stdout=&msg
	if err:=cmd.Run();err!=nil{
		log.Println(err)
	}
	log.Println(msg.String())
}


//3.OutPut 输出|输出合并
func RunDemoOutput(){
	cmd:=exec.Command("ls")
	b,_:=cmd.Output() //直接获得StdOut的字节信息
	log.Println(string(b))

	output,_:=cmd.CombinedOutput() //将Stdout和Stderr合并后输出字节信息
	log.Println(string(output))
}

//4.各个参数
func RunDemoOther(){
	cmd:=exec.Command("ls")
	cmd.Run()
	log.Println(cmd.Process.Pid) //获取当前进程的PID
	log.Println(cmd.ProcessState) //获取结束的标志 正常结束为0
	log.Println(cmd.Path) //命令执行文件的路径
	log.Println(cmd.Args) //命令执行的参数
}

func main() {
	//RunDemo1()
	//RunDemoDir()
	//RunDemoOutput()
	//RunDemoOther()
}
