package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Proxy struct {
}

func (p Proxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transport := http.DefaultTransport
	// step 1. 浅拷贝对象 然后对新对象的属性进行修改 变成新请求发送给服务端
	newReq := new(http.Request)
	*newReq = *req
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := newReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ",") + "," + clientIP
		}
		newReq.Header.Set("X-Forwarded-For", clientIP)
		fmt.Println("clientIP:", clientIP)
	}
	// step 2. 通过RoundTrip请求下游
	resp, err := transport.RoundTrip(newReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	// step 3. 把信息返回给上游服务
	for k, value := range resp.Header {
		for _, v := range value {
			rw.Header().Add(k, v)
		}
	}
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}
func main() {
	fmt.Println("Serve on 8080")
	http.Handle("/", &Proxy{})
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}
}
