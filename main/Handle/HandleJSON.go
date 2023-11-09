package Handle

import (
	Format "GoRuleMatch/main/AllFormat"
	"encoding/json"
	"fmt"
	"os"
)

var pocStruct Format.PocStruct

func TryToParsePocStruct(jsonData string) Format.PocStruct {
	err := json.Unmarshal([]byte(jsonData), &pocStruct)
	if err != nil {
		fmt.Println("[-] Error unmarshal Json:", err)
		os.Exit(1)
	}
	return pocStruct
}
