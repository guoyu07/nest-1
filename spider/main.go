package main

import (
	"github.com/gocolly/colly"
	"log"
	"sync"
	"time"
	"io/ioutil"
	"path/filepath"
)

const (
	AIM = "ershoufang"
	SLAVE_NUM = 10
)

var (
	urlStack = make([]string, 0)
	urlStackM = &sync.Mutex{}
	collector *colly.Collector
	wg = &sync.WaitGroup{}
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	collector = colly.NewCollector(
		colly.AllowedDomains("sh.lianjia.com"),
	)

	collector.UserAgent = "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"

	collector.OnError(func(r *colly.Response, err error) {
		log.Println("[error]", r.Request.URL.String(), err)
	})

	collector.OnResponse(func(r *colly.Response) {
		if March(r.Request.URL.Path) {
			log.Println("find", r.Request.URL)
		} else {
			log.Println("visit", r.Request.URL)
		}
		if IsAim(r.Request.URL.Path) {
			if err := ioutil.WriteFile("data/" + filepath.Base(r.Request.URL.Path), r.Body, 0666); err != nil {
				log.Println("[error]", err)
			}
		}
	})

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if March(link) {
			addUrl(e.Request.AbsoluteURL(link))
		}
	})

	go monitor()

	addUrl("http://sh.lianjia.com/about/sitemap.html")
	wg.Add(SLAVE_NUM)
	for i := 0; i < SLAVE_NUM; i++{
		go slave()
	}
	wg.Wait()
}

func slave() {
	defer wg.Done()
	emptyCount := 0
	for true {
		url := getUrl()
		if url == "" {
			if emptyCount >= 60 {
				return
			} else {
				emptyCount++
				time.Sleep(time.Second)
				continue
			}
		}
		collector.Visit(url)
	}
}

func getUrl() string {
	urlStackM.Lock()
	defer urlStackM.Unlock()
	if len(urlStack) == 0 {
		return ""
	}
	url := urlStack[len(urlStack) - 1]
	urlStack = urlStack[:len(urlStack) - 1]
	return url
}

func addUrl(url string) {
	urlStackM.Lock()
	defer urlStackM.Unlock()
	urlStack = append(urlStack, url)
}

func monitor() {
	for true {
		log.Printf("url stack: len %v, cap %v", len(urlStack), cap(urlStack))
		time.Sleep(10 * time.Second)
	}
}


