package parser

import (
	"awesomeProject/crawler/types"
	"awesomeProject/crawler/za/model"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

var objectInfoRe = regexp.MustCompile(`("objectInfo": {[\s\S]*)"interest"`)

func ParseUser(content []byte) types.ParseResult {
	profile := model.Profile{}
	objectInfoStr := strings.TrimSpace(extractString(content, objectInfoRe))
	objectInfoStr = objectInfoStr[:strings.LastIndex(objectInfoStr, ",")]
	objectInfoStr = fmt.Sprintf("{%s}", objectInfoStr)
	var objectInfo map[string]interface{}
	err := json.Unmarshal([]byte(objectInfoStr), &objectInfo)
	if err == nil {
		fmt.Println(objectInfo)
		// 拼接信息
		//profile.Name = objectInfo["nickname"].(string)
		//profile.Car = objectInfo["momentCount"].(string)
		//profile.Height = objectInfo["heightString"].(string)
		//profile.Weight = objectInfo["momentCount"].(string)
		//profile.Income = objectInfo["salaryString"].(string)
		//profile.House = objectInfo["momentCount"].(string)
		//profile.HuKou = objectInfo["momentCount"].(string)
		//profile.XinZuo = objectInfo["momentCount"].(string)
		//profile.Age = objectInfo["age"].(int)
		//profile.Education = objectInfo["educationString"].(string)
		//profile.Occupation = objectInfo["momentCount"].(string)
		//profile.Gender = objectInfo["genderString"].(string)
		//profile.Marriage = objectInfo["marriageString"].(string)
	}
	// 结束递归
	userResult := types.ParseResult{
		Items: []interface{}{profile},
	}
	fmt.Println("Get Item ", profile)
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
