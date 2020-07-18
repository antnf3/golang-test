package main

import (
	"errors"
	"fmt"
	"time"
)

var errRequestFaield = errors.New("Request Failed")

// func main() {
// 	var results = make(map[string]string)
// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://www.google.com/",
// 		"https://soundcloud.com/",
// 		"https://www.facebook.com/",
// 		"https://www.instagram.com/",
// 		"http://academy.nomadcoders.co/",
// 	}
// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)
// 		if err != nil {
// 			result = "FAILED"
// 		}
// 		results[url] = result
// 	}
// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}
// }

// func hitURL(url string) error {
// 	fmt.Println("Checking: ", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err, resp.Status)
// 		return errRequestFaield
// 	}
// 	return nil
// }

func main() {
	c := make(chan string)
	people := [2]string{"min", "hyun"}

	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println("Waiting message")
	resultOne := <-c
	resultTwo := <-c
	fmt.Println(resultOne)
	fmt.Println(resultTwo)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
