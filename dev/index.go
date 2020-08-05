
package main

import ("net/http")
import "time"
import "io/ioutil"
import "fmt"


func fetchData(c chan string){
	resp, err := http.Get("http://codemobiles.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	c<-string(body)
}

func main() {
	c := make(chan string)
	go fetchData(c)

	for {
		select{
		   case <-c:
			msg := <-c
			fmt.Println(msg)
		   default:
			 println("Waiting for data")
			 time.Sleep(1 * time.Second)
		}
	 }

}