package main

import (
	"fmt"
	"rocketmqSync/cmd"
)

func init() {
	var userinfo = make(map[string]string)
	userinfo["username"] = "1111"
	userinfo["auther"] = "songjincheng"
	fmt.Println(userinfo)

}

func main() {
	cmd.Execute()
}
