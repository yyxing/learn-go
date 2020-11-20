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
	result := ParseUser(file)
	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:      "林姑娘",
		Age:       34,
		Height:    "173cm",
		Weight:    "65kg",
		Income:    "5001-8000元",
		Gender:    "女士",
		XinZuo:    "天秤座(09.23-10.22)",
		Marriage:  "离异",
		Education: "大专",
		HuKou:     "江苏淮安",
		House:     "打算婚后购房",
		Car:       "未买车",
	}
	if profile != expected {
		t.Errorf("expected %v\n, but was %v\n", expected, profile)
	}
}
