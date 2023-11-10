package GlobalVar

import Format "GoRuleMatch/main/AllFormat"

var PocStruct Format.PocStruct

// SetGlobalVariable 设置全局变量的函数
func SetGlobalVariable(value Format.PocStruct) {
	PocStruct = value
}
