package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	t := time.Now()
	tt := time.Now()
	fmt.Println("let's go:", t)
	url := "http://192.168.2.100:10000/api/v/test/id/1?"

	for i := 990; i < 1000; i++ {
		res, err := http.Get(fmt.Sprintf("%s%v", url, i))
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(i, res.StatusCode, time.Since(tt).Milliseconds())
		tt = time.Now()
		// fmt.Printf("%v:%v\n", i, res.StatusCode)
	}

	fmt.Println("go done:", time.Since(t).Milliseconds())

}
