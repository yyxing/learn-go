package filelist

import (
	"awesomeProject/func/fib"
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func CreateFile(filename string) {
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		if os.IsNotExist(err); len(filename) == 0 || filename == "" {
			os.Create(filename)
		}
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

type BaseError string

func (e BaseError) Error() string {
	return e.Message()
}

func (e BaseError) Message() string {
	return string(e)
}

const pathPrefix = "/list/"

func HandleFileList(writer http.ResponseWriter,
	request *http.Request) error {
	if strings.Index(request.URL.Path, pathPrefix) != 0 {
		return BaseError(
			fmt.Sprintf("path %s must start "+
				"with %s",
				request.URL.Path, pathPrefix))
	}
	path := request.URL.Path[len(pathPrefix):]
	if path == "" || len(path) == 0 {
		return BaseError(
			fmt.Sprintf("filename must not empty"),
		)
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(bytes)
	return nil
}
