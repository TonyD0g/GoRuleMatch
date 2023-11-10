package main

import (
	"GoRuleMatch/main/Core"
	"GoRuleMatch/main/GlobalVar"
	"GoRuleMatch/main/Handle"
	"GoRuleMatch/main/Input"
	"GoRuleMatch/main/Log"
	Core2 "GoRuleMatch/main/ShermieProxy/Core"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("              _____       _      __  __       _       _     \n             |  __ \\     | |    |  \\/  |     | |     | |    \n   __ _  ___ | |__) |   _| | ___| \\  / | __ _| |_ ___| |__  \n  / _` |/ _ \\|  _  / | | | |/ _ \\ |\\/| |/ _` | __/ __| '_ \\ \n | (_| | (_) | | \\ \\ |_| | |  __/ |  | | (_| | || (__| | | |\n  \\__, |\\___/|_|  \\_\\__,_|_|\\___|_|  |_|\\__,_|\\__\\___|_| |_|\n   __/ |                                                    \n  |___/                                                     ")
	fmt.Println("基于 Json 、自定义Go脚本的规则匹配验证，可适用于验证本机http流量中是否有符合规则的情况。\nVersion 1.2")

	args := os.Args
	if len(args) == 1 {
		fmt.Println("使用说明:	-ini C:/config.ini\nconfig.ini内容如下:")
	} else if args[1] != "-ini" {
		fmt.Println("[-] 参数错误,例子:")
	}
	inputIniFile := flag.String("ini", ".\\config.ini", "Input the ini file")
	flag.Parse()
	config := Input.HandleIni(*inputIniFile)

	// Determine whether the number of parameters is correct
	if !strings.Contains(config["vul"], ".json") {
		Log.Log.Fatal("[-] 参数错误,例子:")
	}

	Log.NewLogger().Init()
	err := Core2.NewCertificate().Init(config["certFile"], config["keyFile"])
	if err != nil {
		Log.Log.Fatal("[-] 初始化证书失败,请检测证书是否在正确的目录下或者证书是否正确!")
		return
	}

	Log.Log.Println("[+] 规则匹配器已开启,会根据自定义规则将匹配成功的url输出出来 :")
	nagleValue, err := strconv.ParseBool(config["nagle"])
	if err != nil {
		Log.Log.Fatal("[-] nagle 非 bool 值,请重新输入")
	}
	GlobalVar.PocStruct = Handle.TryToParsePocStruct(config["vul"])
	Core.ProxyMain(config["port"], config["network"], config["to"], config["proxy"], nagleValue)
	Log.Log.Println("\n[+] 规则匹配器已关闭")
}
