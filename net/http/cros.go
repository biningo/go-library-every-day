package main

import "net/http"

/**
*@Author icepan
*@Date 2020/8/6
*
**/
//处理跨域请求

func cors(f http.HandlerFunc) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		writer.Header().Add("Access-Control-Allow-Credentials", "true") //设置为true，允许ajax异步请求带cookie信息
		writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
		writer.Header().Set("content-type", "application/json;charset=UTF-8")             //返回数据格式是json
		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusNoContent)
			return
		}
		f(writer,request)
	}
}

func main() {

	http.HandleFunc("/ping", cors(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	}))



	http.ListenAndServe(":9000",nil)

}