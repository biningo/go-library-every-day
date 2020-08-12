package main

import (
	"net/http"
	"net/rpc"
)

/**
*@Author icepan
*@Date 2020/8/12 下午10:13
*@Describe
**/

type NumRequest struct {
	First  int
	Second int
}
type NumResponse struct {
	Result int
}

type Num struct{}

//一定只能返回一个error
func (num *Num) Add(request NumRequest, response *NumResponse) error {
	response.Result = request.First + request.Second
	return nil
}

func main() {
	rpc.Register(&Num{})
	rpc.HandleHTTP()
	http.ListenAndServe(":9000", nil)
}
