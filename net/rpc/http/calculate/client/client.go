package main

import (
	"log"
	"net/rpc"
)

/**
*@Author icepan
*@Date 2020/8/12 下午10:14
*@Describe
**/

type NumRequest struct {
	First int
	Second int
}
type NumResponse struct {
	Result int
}

func main() {
	client,_:=rpc.DialHTTP("tcp","localhost:9000")
	request:=&NumRequest{
		First:  10,
		Second: 5,
	}
	response:=&NumResponse{}

	if err:=client.Call("Num.Add",request,response);err!=nil{
		log.Println(err)
		return
	}
	log.Println(request,response.Result)


}