package fetcher

// fetcher提取器将url的信息 获取返回 传给解析器去解析成需要的数据
import (
	"awesomeProject/crawler/types"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func assembleRequest(request types.Request) (*http.Response, error) {
	req, err := http.NewRequest(request.Method, request.Url, bytes.NewReader(request.Body))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	client := &http.Client{}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36")
	req.Header.Set("Cookie",
		"sid=a9c082ca-c626-4341-95ca-e29dc17e72d6; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1605773095; ec=zbxvWKO1-1605777179320-bd1931af3dd7c1108731840; FSSBBIl1UgzbN7NO=5mzMZLaz_RcpDWNhJMW0_7EZsXvHIpaYmmXG8scJsNtwN_Rgkx3zQ339IYBZtzXqgRhBLTm9x15TAfKlXhKfq3A; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1605840137; FSSBBIl1UgzbN7NP=5U4Qzc256hA9qqqmTggI0KGd2tA3y5I8l9ffKcq3NArXHU9TnwUoTf5qg8brojefUWxdzbii4yYQpqCiXEQqbJI_WkTXjMu0HxPsuqHW_yZZvCPrWmHW55BswAXJal5q0rDfhkxvu5Wlk.u1bx1ry0To3nHu0qoVW1awhAp4VPhBm6RqXkf.FkC51O0pMMrDJVWs6dudUxTKUUwdzdDXGjYKJfOzWIBwMZgbEK2XLwcjU7T5QttOsw4RgkZWin6Tc9VhoxZvR7N.yRZ1dKUFeftD7Z6y7PIKZLk7dOQ3ADP2TO.5ZcFH5p6_M2Oroj5oKQ; _efmdata=yqa8RHsL4kEFGP29RTXfIpwCLmf8zTlp52ilYUtBTG7jurQgez1l2Pbhglm2pP96AHVymMZ37q6pdu6r8zYYgfG7eylWIQNv2g8kKV8wbB4%3D; _exid=A0Y0pYfLtxSgRD3t5Auw47OMb3gSPRe2wMfiH6uJUEGoH5upEYNvT2wYORMUxtWi8%2F5%2FRhUIZsoT%2BwfGwAY62g%3D%3D")
	req.Header.Set("Origin", request.Url)
	req.Header.Set("Referer", request.Url)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Content-Length", "2")
	return client.Do(req)
}

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetcher(request types.Request) ([]byte, error) {
	<-rateLimiter
	resp, err := assembleRequest(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("fetching url ", request.Url)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: response message is %s", resp.Body)
	}
	respReader := bufio.NewReader(resp.Body)
	e := determineEncoding(respReader)
	return ioutil.ReadAll(transform.NewReader(respReader, e.NewDecoder()))
}

func determineEncoding(r io.Reader) encoding.Encoding {
	data, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(data, "")
	return e
}
