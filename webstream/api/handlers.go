package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/kataras/iris/v12"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"webstream/api/entity"
	"webstream/api/infrastructure/mapper"
)

func CreateUser(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		handleError(ServerError, writer)
		return
	}
	defer request.Body.Close()
	var user entity.User
	err = json.Unmarshal(req, &user)
	if err != nil {
		handleError(RequestBodyParseFailedError, writer)
		return
	}
	isExist, err := mapper.FindUserNameIsExist(user.Username)
	if err != nil {
		handleError(ServerError, writer)
		return
	}
	if isExist {
		sendResponse(writer, Fail(HttpBadRequest, "用户名已存在", nil))
		return
	}
	createUser, err := mapper.CreateUser(&user)
	if err != nil {
		handleError(ServerError, writer)
		return
	}
	log.Printf("insert user success %v", createUser)
	sendResponse(writer, Success("", createUser))
}
func QueryUserInfo(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	uname := p.ByName("username")
	userInfo, err := mapper.FindUserByUserName(uname)
	if err != nil {
		handleError(ServerError, writer)
		return
	}
	if userInfo == nil {
		sendResponse(writer, Fail(HttpNotFound, "该用户不存在！", nil))
		return
	}
	sendResponse(writer, Success("", userInfo))
}
func UpdateUser(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

}

func DeleteUser(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {

}

type Person struct {
	Name string
	Age  string
}

func TestQps(c *gin.Context) {
	c.JSON(http.StatusOK, Person{Name: "hello", Age: "world"})
}
func TestQpsIris(c iris.Context) {
	c.JSON(Person{Name: "hello", Age: "world"})
}

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func DoJSONWrite(ctx *routing.Context) error {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(http.StatusOK)
	start := time.Now()
	if err := json.NewEncoder(ctx).Encode(Person{Name: "hello", Age: "world"}); err != nil {
		elapsed := time.Since(start)
		fmt.Print("", elapsed, err.Error(), Person{Name: "hello", Age: "world"})
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return err
	}
	return nil
}

func Login(writer http.ResponseWriter, request *http.Request, p httprouter.Params) {
	req, err := ioutil.ReadAll(request.Body)
	if err != nil {
		handleError(RequestBodyParseFailedError, writer)
		return
	}
	defer request.Body.Close()
	var user *entity.User
	err = json.Unmarshal(req, &user)
	if err != nil {
		handleError(RequestBodyParseFailedError, writer)
		return
	}
	result, err := mapper.GetUserCredential(user.Username, user.Password)
	if err != nil {
		handleError(ServerError, writer)
		return
	}
	if !result {
		sendResponse(writer, Fail(HttpUnauthorized, "账号或者密码错误", nil))
		return
	}
	sendResponse(writer, Success("登陆成功", user.Username))
}
