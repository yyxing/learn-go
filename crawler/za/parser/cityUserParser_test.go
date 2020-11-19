package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityUser(t *testing.T) {
	file, err := ioutil.ReadFile("city_user_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityUser(file)
	if len(result.Items) <= 0 {
		t.Errorf("expect user len > 0, but is %d", len(result.Items))
	}
}
