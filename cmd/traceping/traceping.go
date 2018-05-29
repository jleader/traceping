package main

import (
	"fmt"

	"github.com/lixiangzhong/traceroute"
)

func main() {
	t := traceroute.New("qq.com")
	result, err := t.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
