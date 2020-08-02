package main

import (
	"log"
	"sync"
)

/**
*@Author bingo
*@Date 2020/8/1
*
**/

//sync.Pool
//并发安全
func Demo1() {
	type Stu struct {
		Name string
		Age  int16
	}

	stuPool := sync.Pool{}
	stuPool.New = func() interface{} { //当pool没有对象时调用此方法返回
		return Stu{Name: "None"}
	}

	s1 := Stu{
		Name: "s1",
		Age:  1,
	}
	s2 := Stu{
		Name: "s2",
		Age:  2,
	}
	s3 := Stu{
		Name: "s3",
		Age:  3,
	}

	//放入对象池
	stuPool.Put(s1)
	stuPool.Put(s2)
	stuPool.Put(s3)

	s1 = stuPool.Get().(Stu)
	s2 = stuPool.Get().(Stu)
	s3 = stuPool.Get().(Stu)
	log.Println(s1.Name, s2.Name, s3.Name)
	s4 := stuPool.Get().(Stu) //取完了 则返回默认创建的对象 之后再取的话则创建新对象
	log.Println(s4.Name)

	s4.Name = "s4"
	stuPool.Put(s4)

	s4 = stuPool.Get().(Stu)
	log.Println(s4.Name) //s4

}

func main() {
	Demo1()
}
