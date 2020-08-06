package main

import "net/http"

/**
*@Author icepan
*@Date 2020/8/6
*
**/

func main() {

	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})
	http.ListenAndServe(":9000", nil)

}
