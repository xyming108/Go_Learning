//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	flag.Parse()
	reqStr := strings.Join(flag.Args(), "|")
	fmt.Printf("reqStr = %s\n", reqStr)
	httpGet(reqStr)
}
func httpGet(reqStr string) {
	resp, err := http.Get("http://localhost:8080?" + reqStr)
	if err != nil {
		// 处理错误
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
	}
	fmt.Println(string(body))
}
