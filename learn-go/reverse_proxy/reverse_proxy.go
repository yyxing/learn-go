package main

import (
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
	url1, err := url.Parse(rs1)
	if err != nil {
		log.Fatal(err)
	}
	proxy := NewSingleHostReverseProxy(url1)
	log.Println("Starting http server at", addr)
	log.Println(http.ListenAndServe(addr, proxy))
}
func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
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
			newPayload := []byte("hello " + string(oldPayload))
			nopCloser := ioutil.NopCloser(bytes.NewBuffer(newPayload))
			response.Body = nopCloser
			response.ContentLength = int64(len(newPayload))
			response.Header.Set("Content-Length", fmt.Sprint(len(newPayload)))
		}
		return nil
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc}
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
