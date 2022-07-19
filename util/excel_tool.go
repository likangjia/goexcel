package util

import (
	"regexp"
	"strings"
)

/**
 * @Author: likangjia
 * @Description:
 * @File: excel_tool
 * @Version: 1.0.0
 * @Date: 2022/7/15 12:07 下午
 */
var variableKeyFormat = "^{.*}$"

func GetReplaceKey(key []string) []string {
	for k, v := range key {
		key[k] = machKey(v)
	}
	return key
}

// 判断是否是变量
func machKey(key string) string {
	matched, _ := regexp.MatchString(variableKeyFormat, key)
	if matched{
		key = strings.Trim(key, "{")
		key =  strings.Trim(key, "}")
	}
	return key
}
