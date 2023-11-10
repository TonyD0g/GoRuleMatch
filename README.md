# GoRuleMatch

基于 Json 的规则匹配验证，可适用于验证本机http流量中是否有符合规则的情况。

使用教程：

- 下载 [Releases](https://github.com/TonyD0g/GoRuleMatch/releases) 或下载源码使用ide自行编译

- 创建一个 config.ini 文件

  内容如下：

  ```md
  -vul	// (必须) json 规则
  D:\Coding\Github\GoRuleMatch\1.json
  -port	// (必须) 代理端口
  8080
  -nagle	// (必须) 是否开启nagle数据合并算法
  true
  -certFile // (必须) 证书文件地址,不指定会自动生成
  ./cert.crt
  -keyFile // (必须) 证书文件地址,不指定会自动生成
  ./cert.key
  -to	// (非必须) 代理tcp服务时,目的服务器的ip和端口,默认为0,仅tcp代理使用
  -proxy // (非必须) 上级代理地址
  ```

- 浏览器使用代理插件开启指定代理，比如 8080

- goRuleMatch.exe -ini D:/config.ini 进行执行

  