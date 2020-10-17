package main

import (
	"log"
	"sync"
)

/**
*@Author icepan
*@Date 2020/10/17 上午10:58
*@Describe
**/

func M1() {
	m := sync.Map{}
	m.Store("one", "a")
	v, ok := m.Load("one")
	if !ok {
		return
	}
	log.Println(v)
	m.Delete("one")
	v, ok = m.LoadOrStore("tow", "b")
	if !ok {
		v, ok = m.Load("tow")
		log.Println(v)
		return
	}
	log.Println(v)
}

func main() {
	M1()
}
