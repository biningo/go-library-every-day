package main

import (
	"net/http"
)

/**
*@Author icepan
*@Date 2020/8/6
*
**/

//继承Handle接口
type MyFunc func(http.ResponseWriter, *http.Request)

func (f MyFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main() {

	ping1 := func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ping1"))
	}


	//处理路由功能 所有请求都发往 serverMux的ServerHTTP方法 然后进行路由 分发到具体的Handle
	serverMux := http.NewServeMux()
	serverMux.Handle("/ping1", MyFunc(ping1)) //强转
	//最终还是调用Handle  如上一样强转为 HandleFunc
	serverMux.HandleFunc("/ping2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ping2"))
	})

	//创建server 启动tcp监听
	server := http.Server{Addr: ":9000", Handler: serverMux}
	server.ListenAndServe()

}
