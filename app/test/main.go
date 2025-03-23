package main

import (
	"fmt"
	"github.com/zhangpetergo/LiveStreamRecorder/app/resolver/douyin"
)

func main() {
	url := "https://live.douyin.com/269218797829"

	data, err := douyin.GetStreamData(url)
	if err != nil {
		fmt.Println("GetStreamData error:", err)
		return
	}
	fmt.Println(data)

}
