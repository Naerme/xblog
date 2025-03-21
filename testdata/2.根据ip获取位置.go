// testdata/10.获取地理位置.go
package main

import (
	"blogx_server/core"
	"fmt"
)

func main() {
	core.InitIPDB()

	fmt.Println(core.GetIpAddr("175.0.201.207"))
	fmt.Println(core.GetIpAddr("10.0.201.207"))
	fmt.Println(core.GetIpAddr("110.0.25.207"))
}
