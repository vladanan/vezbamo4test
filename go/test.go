package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type calls struct {
	call int
	mu   sync.Mutex
}

func (cc *calls) caller(c chan int, i int, cr int) {

	// ti := time.Now()
	url := "http://192.168.0.226:10000/api/v/test/id/1?"
	// url = "http://127.0.0.1:7331//api/v/test/id/1?"
	url = "http://127.0.0.1:10000/custom_apis?"

	_, err := http.Get(fmt.Sprintf("%scall=%v", url, i))
	if err != nil {
		fmt.Print(err)
	}
	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println(cc.call, i, res.StatusCode, time.Since(ti).Milliseconds(), time.Now())
	// } else {
	// 	fmt.Println(cr, cc.call, i)
	// }
	// fmt.Println(res.Status, cc.call, i)
	// ti = time.Now()
	if cc.call == cr-1 {
		fmt.Println("over")
		c <- 0
	}
	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.call = cc.call + 1

	// c <- res.StatusCode

}

func main() {
	ts := time.Now()
	fmt.Println("let's go:", ts)

	c := make(chan int, 100000)
	calls := calls{call: 0}

	for i := 0; i < cap(c); i++ {
		go calls.caller(c, i, cap(c))
	}

	// for statusCode := range c {
	// 	fmt.Println("status code:", statusCode)
	// } //114

	<-c

	fmt.Println("go done:", time.Since(ts).Milliseconds(), time.Now())

}
