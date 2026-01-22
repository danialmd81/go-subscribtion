package main

import (
	"github.com/danialmd81/my-subscribtion/all"
	"github.com/danialmd81/my-subscribtion/subs"
	"github.com/danialmd81/my-subscribtion/telegram"
)

func main() {
	telegram.Run()
	subs.Run()
	all.Run()
	// fmt.Println(surfboard.GenerateSSConfig("all/ss.txt", "surfboard/ss.conf"))
	// fmt.Println(surfboard.GenerateVMessConfig("all/vmess.txt", "surfboard/vmess.conf"))
	// fmt.Println(surfboard.GenerateHysteriaConfig("all/hysteria.txt", "surfboard/hysteria.conf"))

}
