package parser

import (
	"awesomeProject/crawler/za/model"
	"io/ioutil"
	"testing"
)

func TestParseUser(t *testing.T) {
	file, err := ioutil.ReadFile("user_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseUser(file, "黄小仙")
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:       "黄小仙",
		Age:        34,
		Height:     "173CM",
		Weight:     "65KG",
		Income:     "5-8千",
		Gender:     "女",
		XinZuo:     "天秤座(09.23-10.22)",
		Marriage:   "离异",
		Education:  "大专",
		Occupation: "经销商",
		HuKou:      "江苏淮安",
		House:      "未买车",
		Car:        "未购车",
	}
	if profile != expected {
		t.Errorf("expected %v\n, but was %v\n", expected, profile)
	}
}
