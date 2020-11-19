package parser

import (
	"awesomeProject/crawler/types"
	"awesomeProject/crawler/za/model"
	"fmt"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+CM)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+KG)</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
var objectInfoRe = regexp.MustCompile(`"objectInfo":([^?!},]+)`)

func ParseUser(content []byte, name string) types.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(content, ageRe))
	if err == nil {
		profile.Age = age
	}
	fmt.Println(extractString(content, objectInfoRe))
	// 拼接信息
	profile.Car = extractString(content, carRe)
	profile.Height = extractString(content, heightRe)
	profile.Weight = extractString(content, weightRe)
	profile.Income = extractString(content, incomeRe)
	profile.House = extractString(content, houseRe)
	profile.HuKou = extractString(content, hukouRe)
	profile.XinZuo = extractString(content, xinzuoRe)
	profile.Education = extractString(content, educationRe)
	profile.Occupation = extractString(content, occupationRe)
	profile.Gender = extractString(content, genderRe)
	profile.Marriage = extractString(content, marriageRe)
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
