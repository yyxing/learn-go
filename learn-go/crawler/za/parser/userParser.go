package parser

import (
	"awesomeProject/crawler/types"
	"awesomeProject/crawler/za/model"
	"fmt"
	"regexp"
	"strings"

	"github.com/tidwall/gjson"
)

var objectInfoRe = regexp.MustCompile(`{("objectInfo":{[\s\S]*)"interest"`)

func ParseUser(content []byte) types.ParseResult {
	profile := model.Profile{}
	objectInfoStr := extractString(content, objectInfoRe)
	objectInfoStr = objectInfoStr[:strings.LastIndex(objectInfoStr, ",")]
	objectInfoStr = fmt.Sprintf("{%s}", objectInfoStr)
	for _, item := range gjson.Get(objectInfoStr, "objectInfo.basicInfo").Array() {
		if strings.Index(item.String(), "座") != -1 {
			profile.XinZuo = item.String()
		}
		if strings.Index(item.String(), "kg") != -1 {
			profile.Weight = item.String()
		}
	}
	for _, item := range gjson.Get(objectInfoStr, "objectInfo.detailInfo").Array() {
		if strings.Index(item.String(), "车") != -1 {
			profile.Car = item.String()
		}
		if strings.LastIndex(item.String(), "籍贯:") != -1 {
			profile.HuKou = item.String()[strings.LastIndex(item.String(), "籍贯:")+len("籍贯:"):]
		}
		if strings.Index(item.String(), "房") != -1 || strings.Index(item.String(), "住") != -1 {
			profile.House = item.String()
		}
	}
	// 拼接信息
	profile.Name = gjson.Get(objectInfoStr, "objectInfo.nickname").String()
	profile.Height = gjson.Get(objectInfoStr, "objectInfo.heightString").String()
	profile.Income = gjson.Get(objectInfoStr, "objectInfo.salaryString").String()
	profile.Age = gjson.Get(objectInfoStr, "objectInfo.age").Int()
	profile.Education = gjson.Get(objectInfoStr, "objectInfo.educationString").String()
	profile.Gender = gjson.Get(objectInfoStr, "objectInfo.genderString").String()
	profile.Marriage = gjson.Get(objectInfoStr, "objectInfo.marriageString").String()
	// 结束递归
	userResult := types.ParseResult{
		Items: []interface{}{profile},
	}
	return userResult
}

func extractString(data []byte, compile *regexp.Regexp) string {
	match := compile.FindSubmatch(data)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
