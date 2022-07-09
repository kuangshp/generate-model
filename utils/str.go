package utils

import (
	"strings"
)

// Case2Camel 下划线转驼峰(大驼峰)
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1) // 根据_来替换成
	name = strings.Title(name)                 // 全部大写
	return strings.Replace(name, " ", "", -1)  // 删除空格
}

// LowerCamelCase 转换为小驼峰
func LowerCamelCase(name string) string {
	name = Case2Camel(name)
	return strings.ToLower(name[:1]) + name[1:]
}
