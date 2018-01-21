package spider

import (
	"github.com/wolfogre/nest/internal/service/entity"
	"time"
	"strings"
	"fmt"
	"log"
	"github.com/wolfogre/nest/internal/service/util/timeformat"
)

const (
	workTime = 0 // 0 点起来爬
)

var (
	statusMsg string
	statusErr error
)

func StartDaemon() {
	go daemon()
}

func Status() (string, error) {
	return statusMsg, statusErr
}

func daemon() {
	for true {
		log.Println("spider wake up")
		urlTemplate := "https://sh.focus.cn/loupan/j0-300_p$index$/?priceStatus=2&saleStatus=6"
		log.Printf("spider start crawl %v\n", urlTemplate)
		crawl(urlTemplate)

		now := time.Now()
		aim := time.Date(now.Year(), now.Month(), now.Day(), workTime, 0, 0, 0, now.Location())
		if aim.Before(now) {
			aim = aim.Add(time.Hour * 24)
		}
		duration := aim.Sub(now)
		log.Printf("spider sleep %v to %v\n", duration, aim)
		time.Sleep(duration)
	}
}

func crawl(urlTemplate string) {
	log.Println("start crawl")
	retry := false
RETRY:
	if retry {
		time.Sleep(5 * time.Minute)
		statusErr = nil
		statusMsg = "ok"
	}
	retry = true
	count, err := GetLoupanCount(formatUrl(urlTemplate, 1))
	if err != nil {
		log.Printf("get loupan count failed: %v, %v\n", err, formatUrl(urlTemplate, 1))
		statusErr = err
		statusMsg = "get loupan count failed"
		goto RETRY
	}

	links := make([]string, 0)
	i := 1
	ls, err := GetLoupanLinks(formatUrl(urlTemplate, 1))
	for err != errNotFound {
		if err == errNotFound {
			break
		}
		if err != nil {
			log.Printf("get loupan link failed: %v, %v\n", err, formatUrl(urlTemplate, i))
			statusErr = err
			statusMsg = "get loupan link failed"
			goto RETRY
		}
		links = append(links, ls...)
		i++
		ls, err = GetLoupanLinks(formatUrl(urlTemplate, i))
	}

	if len(links) != count {
		log.Printf("get wrong number of links, got %v, expert %v\n", len(links), count)
		statusErr = errDefault
		statusMsg = "get wrong number of links"
		goto RETRY
	}

	infos := make([]*entity.Loupan, 0)
	for _, v := range links {
		info, err := GetLoupanInfo(v)
		if err != nil {
			log.Printf("get loupan info failed: %v, %v\n", err, v)
			statusErr = errDefault
			statusMsg = "get loupan info failed"
			goto RETRY
		}
		infos = append(infos, info)
	}

	err = saveJs(infos, fmt.Sprintf("./static/data%v.js", timeformat.FormatDateNow()))
	if err != nil {
		log.Printf("write to file failed: %v\n", err)
		statusErr = errDefault
		statusMsg = "write to file failed"
		goto RETRY
	}

	log.Println("finished crawl")
}

func formatUrl(url string, i int) string {
	return strings.Replace(url, "$index$", fmt.Sprintf("%v", i), 1)
}

