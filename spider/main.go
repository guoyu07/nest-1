package main

import (
	"sync"
	"os"
	"log"
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"path/filepath"
	"flag"
)

const (
	AIM = "ershoufang"
)

var (
	urlCh = make(chan string, 100)
	wg = &sync.WaitGroup{}
	dir = "data"

	start = flag.Int("start", -1, "start index")
	slaveNum = flag.Int("slave", 100, "slave num")
)

func main() {
	flag.Parse()

	if err := os.MkdirAll(dir, os.ModeDir); err != nil && !os.IsExist(err) {
		log.Panic(err)
	}

	wg.Add(*slaveNum)
	for i := 0; i < *slaveNum; i++ {
		go slave()
	}

	if *start == -1 {
		flag.PrintDefaults()
		return
	}

	for i := *start; i < 99999999; i++ {
		if _, err :=  os.Stat(dir + "/sh%d.html"); os.IsNotExist(err) {
			urlCh <- fmt.Sprintf("http://sh.lianjia.com/ershoufang/sh%d.html", i)
		}
	}

	wg.Wait()
}

func slave() {
	defer wg.Done()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for v := range urlCh {
		retry := false
	RETRY:
		if retry {
			time.Sleep(time.Second)
		}
		retry = true
		response, err := client.Get(v)
		if err != nil {
			log.Printf("get %v failed: %v\n", v, err)
			goto RETRY
		}

		if response.StatusCode == 200 {
			buffer, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("read body of %v failed: %v\n", v, err)
				goto RETRY
			}
			err = ioutil.WriteFile(dir + "/" + filepath.Base(v), buffer, 0644)
			if err != nil {
				log.Printf("write body of %v failed: %v\n", v, err)
				goto RETRY
			}
		}
		log.Printf("ok get %v of %v\n", response.StatusCode, v)

		response.Body.Close()
	}
}
