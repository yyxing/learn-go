package main

import (
	"awesomeProject/interface/mock"
	"awesomeProject/interface/real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func inspect(r Retriever) {
	fmt.Printf("%T %v \n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("this is mock.Retriever ", v)
	case *real.Retriever:
		fmt.Println("this is real.Retriever ", v)
	}
}
func main() {
	var r Retriever
	r = &mock.Retriever{Contents: "this is www.baidu.com"}
	inspect(r)

	r = &real.Retriever{UserAgent: "Chrome 5.0", TimeOut: time.Second * 60}
	inspect(r)

	if retriever, ok := r.(*real.Retriever); ok {
		fmt.Println("this is mock.Retriever ", retriever)
	} else {
		fmt.Println("this is not mock.Retriever")
	}
}
