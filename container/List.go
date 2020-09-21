package main

/**
*@Author icepan
*@Date 2020/9/21 下午1:17
*@Describe
**/
import (
	"container/list"
	"log"
)

type Stu struct {
	Name string
}

func L1() {
	l1 := list.New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(Stu{"a"})
	log.Println(l1.Len())
	i := l1.Front().Value.(int)
	s := l1.Back().Value.(Stu)
	log.Println(i, s.Name)
}

func main() {
	L1()
}
