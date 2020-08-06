package main

import "net/http"

/**
*@Author icepan
*@Date 2020/8/6
*
**/

type MyHandler struct {
	Msg string
}

//最终都会执行这个方法来处理请求 也就是继承 Handle接口
func (my MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(my.Msg))
}

func main() {
	http.Handle("/ping", MyHandler{"pong"})

	//第二个参数也是一个Handle 如果不为nil，则替换掉默认的 ServerMux，但是需要自己设置路由，一般都设置为nil
	http.ListenAndServe(":9000", nil)
}
