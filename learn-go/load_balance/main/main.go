package main

import (
	"awesomeProject/load_balance"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var (
	addr = "127.0.0.1:2002"
)

func main() {
	rs1 := "http://127.0.0.1:2003/base"
	rs2 := "http://127.0.0.1:2004/base"
	lb := load_balance.LoadBalanceFactory(load_balance.ConsistentHashLB)
	if err := lb.Add(rs1, "7"); err != nil {
		log.Fatal(err)
	}
	if err := lb.Add(rs2, "3"); err != nil {
		log.Fatal(err)
	}
	proxy := NewMultiHostReverseProxy(lb)
	log.Println("Starting http server at", addr)
	log.Println(http.ListenAndServe(addr, proxy))
}
func NewMultiHostReverseProxy(lb load_balance.LoadBalance) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		nextAddr, err := lb.Get(req.URL.String())
		if err != nil {
			log.Fatal("get next addr fail")
		}
		target, err := url.Parse(nextAddr)
		if err != nil {
			log.Fatal(err)
		}
		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}
	modifyFunc := func(response *http.Response) error {
		if response.StatusCode > 299 && response.StatusCode < 200 {
			// 获取原数据的payload
			oldPayload, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			newPayload := []byte("[ERROR] " + string(oldPayload))
			nopCloser := ioutil.NopCloser(bytes.NewBuffer(newPayload))
			response.Body = nopCloser
			response.ContentLength = int64(len(newPayload))
			response.Header.Set("Content-Length", fmt.Sprint(len(newPayload)))
		} else {
			// 获取原数据的payload
			oldPayload, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			newPayload := []byte("[INFO] " + string(oldPayload))
			nopCloser := ioutil.NopCloser(bytes.NewBuffer(newPayload))
			response.Body = nopCloser
			response.ContentLength = int64(len(newPayload))
			response.Header.Set("Content-Length", fmt.Sprint(len(newPayload)))
		}
		return nil
	}
	errHandler := func(writer http.ResponseWriter, request *http.Request, err error) {

	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc,
		ErrorHandler: errHandler}
}
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}
