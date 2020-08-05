package main

import (
	"errors"
	"log"
)

/**
*@Author icepan
*@Date 2020/8/5
*
**/

//错误处理包

//1. New 创建error
func ErrorsNewDemo() {
	err := errors.New("I'am error") //返回error接口
	log.Println(err.Error())        //返回错误信息:string
}


//2. 自定义错误 继承error接口即可
type StudentError struct {
	msg string
}

//返回错误信息的方法string
func (s StudentError) Error() string {
	return s.msg
}

type Student struct {
	error StudentError //内嵌一个自己的错误
	Name  string
}

func MyErrorDemo(name string) (Student, error) {
	var s Student
	s.Name = name
	if s.Name != "bingo" {
		return s, StudentError{s.Name + " is error"}
	}
	return s, nil
}
func StudentDemo() {
	if s, err := MyErrorDemo("bingo"); err != nil {
		panic(err)
	} else {
		log.Println(s.Name)
	}
}

func main() {

	ErrorsNewDemo()
	StudentDemo()

}
