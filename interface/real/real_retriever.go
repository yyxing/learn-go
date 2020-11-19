package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	get, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	response, err := httputil.DumpResponse(get, true)
	get.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(response)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	post, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	response, err := httputil.DumpResponse(post, true)
	post.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(response)
}
