package main

import (
	"testing"
	"fmt"
)

func TestGetLoupanCount(t *testing.T) {
	fmt.Println(GetLoupanCount("https://sh.focus.cn/loupan/m70-90_w0_p12/?saleStatus=6"))
}

func TestGetLoupanLinks(t *testing.T) {
	fmt.Println(GetLoupanLinks("https://sh.focus.cn/loupan/m70-90_w0_p12/?saleStatus=6"))
	fmt.Println(GetLoupanLinks("https://sh.focus.cn/loupan/m70-90_w0_p120/?saleStatus=6"))
}

func TestGetLoupanInfo(t *testing.T) {
	fmt.Println(GetLoupanInfo("https://sh.focus.cn/loupan/10941.html"))
}