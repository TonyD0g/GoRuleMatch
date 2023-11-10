package Handle

import (
	Format "GoRuleMatch/main/AllFormat"
	"GoRuleMatch/main/Log"
	"encoding/json"
	"io/ioutil"
)

var pocStruct Format.PocStruct

func TryToParsePocStruct(jsonData string) Format.PocStruct {
	// 读取 JSON 文件内容
	jsonDataBytes, err := ioutil.ReadFile(jsonData)
	if err != nil {
		Log.Log.Fatal("[-] Error reading JSON file:", err)
	}
	err = json.Unmarshal(jsonDataBytes, &pocStruct)
	if err != nil {
		Log.Log.Fatal("[-] Error unmarshal Json:", err)
	}
	return pocStruct
}
