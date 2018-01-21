package spider

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"errors"
	"github.com/wolfogre/nest/internal/service/entity"
	"fmt"
	"github.com/wolfogre/nest/internal/service/util/timeformat"
	"strings"
	"time"
)

var (
	errDefault = errors.New("default")
	errNotFound = errors.New("not found")
)

func GetLoupanCount(url string) (int, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return 0, err
	}
	if str := doc.Find(".s-m-fr strong").Text(); str != "" {
		result, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return 0, err
		}
		return int(result), nil
	} else {
		return 0, errNotFound
	}
}

func GetLoupanLinks(url string) ([]string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	if len(doc.Find(".no-result").Nodes) != 0 {
		return nil, errNotFound
	}
	result := make([]string, 0)
	for _, v := range doc.Find(".module-lplist-list .list .txt-center .title a").Nodes {
		node := goquery.NewDocumentFromNode(v)
		if href, ok := node.Attr("href"); ok {
			result = append(result, href)
		}
	}
	return result, nil
}

func GetLoupanInfo(url string) (*entity.Loupan, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	sel := doc.Find(".module-loupan-circum")

	id, ok := sel.Attr("data-group_id")
	if !ok {
		return nil, fmt.Errorf("can not find data-group_id")
	}
	name, ok := sel.Attr("data-house_name")
	if !ok {
		return nil, fmt.Errorf("can not find data-house_name")
	}
	latStr, ok := sel.Attr("data-lat")
	if !ok {
		return nil, fmt.Errorf("can not find data-lat")
	}
	lngStr, ok := sel.Attr("data-lng")
	if !ok {
		return nil, fmt.Errorf("can not find data-lng")
	}
	lat := 0.0
	if latStr != "" {
		lat, err = strconv.ParseFloat(latStr, 64)
		if err != nil {
			return nil, err
		}
	}
	lng := 0.0
	if lngStr != "" {
		lng, err = strconv.ParseFloat(lngStr, 64)
		if err != nil {
			return nil, err
		}
	}

	vaildDateOrigin := doc.Find(".vaild-date").Text()
	if vaildDateOrigin == "" {
		return nil, fmt.Errorf("can not find vaild-date")
	}
	vaildDate := vaildDateOrigin
	vaildDate = strings.TrimPrefix(vaildDate, "(有效期：")
	vaildDate = strings.TrimSuffix(vaildDate, ")")
	vaildSplits := strings.Split(vaildDate, "-")
	if len(vaildSplits) != 2 {
		return nil, fmt.Errorf("wrong vaild date: %v", vaildDateOrigin)
	}
	start := parseVaildDate(vaildSplits[0])
	if start == 0 {
		return nil, fmt.Errorf("wrong vaild date: %v", vaildDateOrigin)
	}
	stop := parseVaildDate(vaildSplits[1])
	if start == 0 {
		return nil, fmt.Errorf("wrong vaild date: %v", vaildDateOrigin)
	}


	return &entity.Loupan{
		Id: id,
		Name: name,
		Url: url,
		Lat: lat,
		Lng: lng,
		StartDate: start,
		StopDate: stop,
		FoundTime: timeformat.FormatTimeNow(),
	}, nil
}

func parseVaildDate(s string) int64 {
	splits := strings.Split(s, ".")
	if len(splits) != 3 {
		return 0
	}
	y, err := strconv.ParseInt(splits[0], 10, 64)
	if err != nil {
		return 0
	}
	y += 2000
	m, err := strconv.ParseInt(splits[1], 10, 64)
	if err != nil {
		return 0
	}
	d, err := strconv.ParseInt(splits[2], 10, 64)
	if err != nil {
		return 0
	}
	date := time.Date(int(y), time.Month(int(m)), int(d), 0, 0, 0, 0, time.Local)
	return timeformat.FormatDate(date)
}