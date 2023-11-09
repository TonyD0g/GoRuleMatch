package main

import (
	"GoRuleMatch/main/Core"
	"GoRuleMatch/main/Input"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" _____      ______      _     ___  ___      _       _     \n|  __ \\     | ___ \\    | |    |  \\/  |     | |     | |    \n| |  \\/ ___ | |_/ /   _| | ___| .  . | __ _| |_ ___| |__  \n| | __ / _ \\|    / | | | |/ _ \\ |\\/| |/ _` | __/ __| '_ \\ \n| |_\\ \\ (_) | |\\ \\ |_| | |  __/ |  | | (_| | || (__| | | |\n \\____/\\___/\\_| \\_\\__,_|_|\\___\\_|  |_/\\__,_|\\__\\___|_| |_|\n                                                          \n                                                          ")
	fmt.Println("基于 Json 、自定义Go脚本的规则匹配验证，可适用于验证本机http流量中是否有符合规则的情况。\nVersion 0.0.1")
	args := os.Args
	if len(args) == 1 {
		fmt.Println("使用说明:	-ini C:/config.ini\nconfig.ini内容如下:\n\n-email // fofa的email (必须)\n-key // fofa的key (必须)\n-url // 扫单个url (非必须)\n-file // 扫url文件中的每一个url (非必须)\n-vul // poc/exp文件,文件后缀为.go (必须)\n-mod // 指定poc/exp这两种模式 (必须)\n-proxy // burpsuite 代理,用于方便写poc/exp (必须)\n-maxConcurrentLevel // 最大并发量,越大扫描速度越快 (必须)\n-maxFofaSize\t   // fofa最大检索数 (必须)")
	} else if args[1] != "-ini" {
		fmt.Println("[-] 参数错误,例子:-email // fofa的email (必须)\n-key // fofa的key (必须)\n-url // 扫单个url (非必须)\n-file // 扫url文件中的每一个url (非必须)\n-vul // poc/exp文件,文件后缀为.go (必须)\n-mod // 指定poc/exp这两种模式 (必须)\n-proxy // burpsuite 代理,用于方便写poc/exp (必须)\n-maxConcurrentLevel // 最大并发量,越大扫描速度越快 (必须)\n-maxFofaSize\t   // fofa最大检索数 (必须)")
		os.Exit(1)
	}
	inputIniFile := flag.String("ini", ".\\config.ini", "Input the ini file")
	flag.Parse()
	config := Input.HandleIni(*inputIniFile)

	// Determine whether the number of parameters is correct
	if !strings.Contains(config["vul"], ".go") {
		fmt.Println("[-] 参数错误,例子:-email // fofa的email (必须)\n-key // fofa的key (必须)\n-url // 扫单个url (非必须)\n-file // 扫url文件中的每一个url (非必须)\n-vul // poc/exp文件,文件后缀为.go (必须)\n-mod // 指定poc/exp这两种模式 (必须)\n-proxy // burpsuite 代理,用于方便写poc/exp (必须)\n-maxConcurrentLevel // 最大并发量,越大扫描速度越快 (必须)\n-maxFofaSize\t   // fofa最大检索数 (必须)")
		os.Exit(1)
	}

	//maxConcurrentLevelInt, err := strconv.Atoi(config["maxConcurrentLevel"])
	//if err != nil {
	//	fmt.Println("The maximum concurrency you entered is not a number!", err)
	//}
	// 两种 poc模式,第一种为json格式,第二种为代码格式
	//var pocStruct format2.PocStruct
	//pocModule := Core.LoadPlugin(config["vul"])
	//if pocModule == 1 {
	//	pocStruct = Handle.TryToParsePocStruct(User.Json)
	//}

	//var urlsList []string
	//fmt.Println("[+] 规则匹配器已开启,会根据自定义规则将匹配成功的url输出出来 :")
	//if pocModule == 1 {
	//	Core.ForSendByJson(urlsList, pocStruct, config["proxy"], maxConcurrentLevelInt)
	//} else {
	//	Core.ForSendByCode(config["mod"], urlsList, config["proxy"], maxConcurrentLevelInt)
	//}
	num, err := strconv.Atoi(config["port"])
	if err != nil {
		fmt.Println("[-] 请检查端口,端口非数字")
		os.Exit(1)
	}
	Core.GetHttpData(num)
	//if Core.IsExploitSuccessByJson(pocStruct, procedureResponse, customRequestBody) {
	//	if splitURL := strings.Split(tmpUrlForAllRequestPath, "?"); len(splitURL) >= 2 {
	//		params := strings.Split(splitURL[1], "&")
	//		encodedParams := make([]string, len(params))
	//		for tmpI := range params {
	//			p := strings.Split(params[tmpI], "=")
	//			encodedParams[tmpI] = url.QueryEscape(p[0]) + "=" + url.QueryEscape(p[1])
	//		}
	//		fmt.Println("[+] [ " + parsedURL.Scheme + "://" + parsedURL.Host + "/" + strings.Join(encodedParams, "&") + " ]\tSuccess! The target may have this vulnerability")
	//	} else {
	//		fmt.Println("[+] [ " + parsedURL.Scheme + "://" + parsedURL.Host + "/" + " ]\tSuccess! The target may have this vulnerability")
	//	}
	//}
	fmt.Println("\n[+] 规则匹配器已关闭")
}
