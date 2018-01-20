package main

import (
	"fmt"
	"strings"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	url := "https://sh.focus.cn/loupan/j100-300_w0_p%index%/?priceStatus=2&saleStatus=6"
	count, err := GetLoupanCount(index(url, 1))
	if err != nil {
		log.Fatal(err)
	}
	links := make([]string, 0)
	i := 1
	ls, err := GetLoupanLinks(index(url, i))
	for err != ErrNotFound {
		log.Println("visit", index(url, i))
		if err == ErrNotFound {
			break
		}
		if err != nil {
			log.Println(index(url, i), err)
		}
		links = append(links, ls...)
		i++
		ls, err = GetLoupanLinks(index(url, i))
	}

	if len(links) != count {
		log.Fatalf("get wrong number of links, got %v, expert %v\n", len(links), count)
	}

	infos := make([]*LoupanInfo, 0)
	for _, v := range links {
		log.Println("visit", v)
		info, err := GetLoupanInfo(v)
		if err != nil {
			log.Println(v, err)
			continue
		}
		infos = append(infos, info)
	}

	for _, v := range infos {
		fmt.Println(v)
	}

}

func index(url string, i int) string {
	return strings.Replace(url, "%index%", fmt.Sprintf("%v", i), 1)
}
