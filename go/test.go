package main
// export PATH=$PATH:/usr/local/go/bin

import (
	"fmt"
	"net/http"
	// "strconv"
	"time"
)

func main() {
	fmt.Println("go let's go")
	t := time.Now()

	url := "http://192.168.2.100:10000/api/v/test/id/1"


	for i := 990; i < 1000; i++ {

		//_, err := http.Get("http://192.168.2.100:10000/custom_apis?"+strconv.Itoa(i))
		res, err := http.Get(fmt.Sprintf("%s?%v", url, i))
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("%v:%v\n", i , res.StatusCode)
	}

	fmt.Println("done!", time.Since(t).Seconds())
}
