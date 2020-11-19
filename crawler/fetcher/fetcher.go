package fetcher

// fetcher提取器将url的信息 获取返回 传给解析器去解析成需要的数据
import (
	"awesomeProject/crawler/types"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetcher(request types.Request) ([]byte, error) {
	resp, err := http.Get(request.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("fetching url ", request.Url)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: response message is %s", resp.Body)
	}
	respReader := bufio.NewReader(resp.Body)
	e := determineEncoding(respReader)
	return ioutil.ReadAll(transform.NewReader(respReader, e.NewDecoder()))
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
