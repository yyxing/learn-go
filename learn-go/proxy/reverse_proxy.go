package main

import (
	"bufio"
	"log"
	"net/http"
	"net/url"
)

var (
	proxyAddr = "http://127.0.0.1:2003"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// step 1. 解析代理地址，并更改请求体的协议和主机
		proxy, err := url.Parse(proxyAddr)
		if err != nil {
			log.Fatal(err)
		}
		request.URL.Host = proxy.Host
		request.URL.Scheme = proxy.Scheme
		// step 2. 请求下游系统
		transport := http.DefaultTransport
		response, err := transport.RoundTrip(request)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		// step 3. 把信息返回给上游服务
		for k, value := range response.Header {
			for _, v := range value {
				writer.Header().Add(k, v)
			}
		}
		_, err = bufio.NewReader(response.Body).WriteTo(writer)
		if err != nil {
			log.Fatal(err)
		}
	})
	err := http.ListenAndServe(":2002", nil)
	if err != nil {
		panic(err)
	}
}
