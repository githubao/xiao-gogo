// main func
// author: baoqiang
// time: 2019-06-03 20:53
package main

import (
	"fmt"
	_ "github.com/githubao/xiao-gogo/gin"
	"github.com/githubao/xiao-gogo/gls"
)

func main() {
	fmt.Println("hello, gogo")
	gls.HelloGls()
	//xretry.HelloRetry()
}

