package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func main() {
	request := gorequest.New()
	resp, body, errs := request.Get("https://www.youtube.com/").End()
	if errs != nil {
		panic(errs)
	}
	if resp.StatusCode == 200 {
		fmt.Println(body[120:140])
	}
}
