package main

import "net/http"

/**
*@Author icepan
*@Date 2020/11/1 上午11:23
*@Describe
**/

func main() {
	c := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	r := &http.Request{}
	c.Do(r)
}
