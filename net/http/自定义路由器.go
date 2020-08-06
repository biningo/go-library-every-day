package main

import (
	"errors"
	"log"
	"net/http"
	"sync"
)

/**
*@Author icepan
*@Date 2020/8/6
*
**/
var (
	PathNotFound                    = errors.New("404 没有page")
	PageNotFoundHandle MyHandleFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("页面没有找到"))
	}
)

type MyMux struct {
	mu        sync.RWMutex
	routerMap map[string]Entry
}
type Entry struct {
	pattern string
	handler http.Handler
}

func (my MyMux) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println(my.routerMap)
	if h, err := my.GetHandler(*request); err != nil {
		PageNotFoundHandle(writer, request)
	} else {
		h.ServeHTTP(writer, request)
	}
}

type MyHandleFunc func(w http.ResponseWriter, r *http.Request)

func (f MyHandleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}
func (my *MyMux) SetHandler(pattern string, handleFunc func(w http.ResponseWriter, r *http.Request)) {
	if my.routerMap==nil{
		my.routerMap = map[string]Entry{}
	}
	e:=Entry{
		pattern: pattern,
		handler: MyHandleFunc(handleFunc),
	}
	my.routerMap[pattern] = e
}
func (my MyMux) GetHandler(r http.Request) (http.Handler, error) {
	path := r.URL.Path
	if entry, flag := my.routerMap[path]; flag == false {
		return nil, PathNotFound
	} else {
		return entry.handler, nil
	}

}

func main() {
	var myMux MyMux

	myMux.SetHandler("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	})
	myMux.SetHandler("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})
	http.ListenAndServe(":8080", myMux)
}
