package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type calls struct {
	call int
	ok   int
	nok  int
	mu   sync.Mutex
}

func syncCaller(url string, cr int) {
	// ti := time.Now()
	ok := 0
	nok := 0

	for i := 0; i < cr; i++ {
		res, err := http.Post(url, "application/json", nil)
		// res, err := http.Get(fmt.Sprintf("%scall=%v", url, i))
		if err != nil {
			log.Println(err)
		}
		if res.StatusCode != http.StatusOK {
			nok++
			// fmt.Println(i, res.StatusCode, time.Since(ti).Milliseconds(), time.Now())
		} else {
			ok++
			// fmt.Println(i)
		}
		fmt.Println(res.Status, i)
		// ti = time.Now()
	}
	fmt.Println("ok:", ok, "nok:", nok)
}

type sign_up_post struct {
	email1    string
	user_name string
	password1 string
}

func (cc *calls) asyncCaller(urlString string, c chan int, i int, cr int) {
	// ti := time.Now()

	u := sign_up_post{
		email1:    "y.emailbox-vezbamo@yahoo.com",
		user_name: "pera5678",
		password1: "pera1234",
	}
	jsonData, err := json.Marshal(u)
	if err != nil {
		log.Println(err, jsonData)
	}

	v := url.Values{}
	// sign_up
	v.Set("email1", "y.emailbox-vezbamo@yahoo.com")
	v.Set("email2", "y.emailbox-vezbamo@yahoo.com")
	v.Set("user_name", "neki-user")
	v.Set("password1", "neki-pass")
	v.Set("password2", "neki-pass")
	// sign_in

	v.Set("email", "vladan_zasve@yahoo.com")
	v.Set("password", "321654987")
	res, err := http.PostForm(urlString, v)
	// res, err := http.Post(urlString, "application/json", nil)
	// res, err := http.Get(fmt.Sprintf("%scall=%v", urlString, i))
	if err != nil {
		fmt.Print(err)
	}
	if res.StatusCode != http.StatusOK {
		cc.nok++
		// fmt.Println(cc.call, i, res.StatusCode, time.Since(ti).Milliseconds(), time.Now())
	} else {
		cc.ok++
		// fmt.Println(cr, cc.call, i)
	}
	fmt.Println(res.Status, cc.call, i)
	// ti = time.Now()
	if cc.call == cr-1 {
		fmt.Println("over, ok:", cc.ok, "nok", cc.nok)
		c <- 0
	}
	cc.mu.Lock()
	defer cc.mu.Unlock()
	cc.call = cc.call + 1
	// c <- res.StatusCode
}

func main() {

	url := "http://192.168.0.226:10000/api/v/test/id/1?"
	url = "http://127.0.0.1:7331/api/v/test/id/1?"
	// url = "http://127.0.0.1:10000/custom_apis?"
	// url = "http://192.168.2.100:7331/custom_apis?"
	// url = "http://127.0.0.1:7331/custom_apis?"
	url = "http://127.0.0.1:7331/auth/sign_up_post"
	// url = "http://192.168.2.100:10000/auth/sign_up_post?email1=y.emailbox-vezbamo@yahoo.com&email2=y.emailbox-vezbamo@yahoo.com&user_name=pera5678&password1=pera1234&password2=pera1234&sign_up=sign_up"
	url = "http://192.168.2.100:10000/auth/sign_up_post"
	url = "http://192.168.2.100:10000/auth/sign_in_post"

	ts := time.Now()
	fmt.Println("let's go:", ts)

	mode := "async"

	c := make(chan int, 3)
	calls := calls{call: 0}

	if mode == "async" {
		for i := 0; i < cap(c); i++ {
			go calls.asyncCaller(url, c, i, cap(c))
		}
		<-c
	} else {
		syncCaller(url, cap(c))
	}

	fmt.Println("go done:", time.Since(ts).Milliseconds(), time.Now())

}
