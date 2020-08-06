package main

import (
	"context"
	"log"
	"time"
)

/**
*@Author icepan
*@Date 2020/8/6
*
**/
func main() {

	p1, _ := context.WithTimeout(context.Background(), time.Second*5)
	p2 := context.WithValue(p1, "age", 18)
	go RunAge(p2)

	time.Sleep(time.Hour)
}

func RunAge(ctx context.Context) {

	p3 := context.WithValue(ctx, "name", "jack")
	go RunName(p3)
	for {

		select {
		case <-ctx.Done():
			log.Println("Age结束")
			return
		default:
			age := ctx.Value("age")
			log.Println("age：",age)
			time.Sleep(time.Second)
		}

	}
}


func RunName(ctx context.Context){
	for {

		select {
		case <-ctx.Done():
			log.Println("Name结束")
			return
		default:
			name := ctx.Value("name")
			log.Println("name：",name)
			time.Sleep(time.Second)
		}

	}


}