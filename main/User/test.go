package User

import (
	Format "GoRuleMatch/main/AllFormat"
	"net/http"
)

var Json string
var Poc func(hostInfo string, client *http.Client) bool
var Exp func(expResult Format.ExpResult, client *http.Client) Format.ExpResult

func init() {
	// 有代码使用代码,无代码使用json
	// 如果存在代码,可以不写Json格式(即Json格式有架构,但内容为空).但必须存在 fofa语句
	// 此处的json只是说明json的使用方式,与代码模式并无关联
	Json = `{
    // 必须,表明想要查找的fofa语句.
    "fofa":"body=\"hello world\"", 
   	// 请求包
    "Request":{
        // 请求方法
		"Method": "GET",
		 // 请求路径,这里分别请求两个uri
		"Uri": [
		      "/robots.txt",
               "/hello.txt"
	   			],
		// 自定义 header 头
		"Header":{
			"Accept-Encoding":"gzip"
		}
	},
    // 响应包
    "Response":{
        // 定义多个Group之间的关系,有AND和OR这两种,其中AND是都满足漏洞才存在,OR是其中一个条件满足即可.
		"Operation":"OR",
        // 判断条件
		"Group":[
            	 // 条件1
				{
                    // 支持正则表达式
                     "Regexp": ".*?",
			        "Header":{
                        				// 状态码
			                            "Status": "200"      
			            },
			        // response Body ,同样是支持多个Body,当都符合时为True
			        "Body":[
			        				    "Hello World",
                        				 "nice"
			            ]  
			    },
            	 // 条件2
           		 {
			        "Header":{
                        				// 状态码
			                            "Status": "200"      
			            }
			    }
			]

}
}`
}
