package context

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type Handler func(*Context)

const (
	// ContentTypeHeaderKey is the header key of "Content-Type".
	ContentTypeHeaderKey = "Content-Type"
	// ContentBinaryHeaderValue header value for binary data.
	ContentBinaryHeaderValue = "application/octet-stream"
	// ContentWebassemblyHeaderValue header value for web assembly files.
	ContentWebassemblyHeaderValue = "application/wasm"
	// ContentHTMLHeaderValue is the  string of text/html response header's content type value.
	ContentHTMLHeaderValue = "text/html"
	// ContentJSONHeaderValue header value for JSON data.
	ContentJSONHeaderValue = "application/json"
	// ContentJSONProblemHeaderValue header value for JSON API problem error.
	// Read more at: https://tools.ietf.org/html/rfc7807
	ContentJSONProblemHeaderValue = "application/problem+json"
	// ContentXMLProblemHeaderValue header value for XML API problem error.
	// Read more at: https://tools.ietf.org/html/rfc7807
	ContentXMLProblemHeaderValue = "application/problem+xml"
	// ContentJavascriptHeaderValue header value for JSONP & Javascript data.
	ContentJavascriptHeaderValue = "application/javascript"
	// ContentTextHeaderValue header value for Text data.
	ContentTextHeaderValue = "text/plain"
	// ContentXMLHeaderValue header value for XML data.
	ContentXMLHeaderValue = "text/xml"
	// ContentXMLUnreadableHeaderValue obselete header value for XML.
	ContentXMLUnreadableHeaderValue = "application/xml"
	// ContentMarkdownHeaderValue custom key/content type, the real is the text/html.
	ContentMarkdownHeaderValue = "text/markdown"
	// ContentYAMLHeaderValue header value for YAML data.
	ContentYAMLHeaderValue = "application/x-yaml"
	// ContentFormHeaderValue header value for post form data.
	ContentFormHeaderValue = "application/x-www-form-urlencoded"
	// ContentFormMultipartHeaderValue header value for post multipart form data.
	ContentFormMultipartHeaderValue = "multipart/form-data"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Path    string
	Method  string
	Params  map[string]string
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Writer:  writer,
		Request: request,
		Path:    request.URL.Path,
		Method:  request.Method,
	}
}

// 获取header中的值
func (c *Context) HeaderValue(key string) string {
	return c.Request.URL.Query().Get(key)
}

// 写入json数据的上层方法
func (c *Context) JSON(v interface{}) (int, error) {
	c.ContentType(ContentJSONHeaderValue)
	return WriterJSON(c.Writer, v)
}

// 写入html数据
func (c *Context) HTML(html string) (int, error) {
	c.ContentType(ContentHTMLHeaderValue)
	return c.Writer.Write([]byte(html))
}

// 写入普通数据
func (c *Context) Write(body []byte) (int, error) {
	c.ContentType(ContentTextHeaderValue)
	return c.Writer.Write(body)
}
func (c *Context) GetParam(key string) string {
	return c.Params[key]
}
func (c *Context) WriteString(str string) (int, error) {
	return c.Write([]byte(str))
}

// 设置状态码
func (c *Context) StatusCode(statusCode int) {
	c.Writer.WriteHeader(statusCode)
}

// 设置ContentType
func (c *Context) ContentType(cType string) {
	c.Header(ContentTypeHeaderKey, cType)
}

// 设置返回请求header
func (c *Context) Header(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 将json数据写入write流
func WriterJSON(writer io.Writer, v interface{}) (int, error) {
	marshal, err := json.Marshal(v)
	if err != nil {
		return 0, err
	}
	return writer.Write(marshal)
}
