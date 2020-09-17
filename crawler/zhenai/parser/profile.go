// 解析用户详情页，提取用户信息
package parser

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/model"

	"github.com/zhibailin/go-distributed-crawler-from-scratch/crawler/engine"
)

var ageRex = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var marriageRex = regexp.MustCompile(`<td><span class="label">婚况：</span>([^>]+)</td>`)

func ParseProfile(contents []byte) engine.ParseResult {

	// 跳过反反爬的处理，除了用户名，其他字段统一用固定值填充
	// 整数型字段的处理
	age, err := strconv.Atoi(extractString(contents, ageRex))
	if err == nil {
		fmt.Println(age)
	}
	// 字符串型字段的处理
	fmt.Println(extractString(contents, marriageRex))

	profile := model.Profile{
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "3001-5000元",
		Gender:     "女",
		Xinzuo:     "牧羊座",
		Occupation: "人事/行政",
		Marriage:   "离异",
		House:      "已够房",
		Hokou:      "山东菏泽",
		Education:  "大学本科",
		Car:        "未购车",
	}

	result := engine.ParseResult{
		Items: []interface{}{profile}, // slice of interface
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
