package main

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
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
		return 0, ErrNotFound
	}
}

func GetLoupanLinks(url string) ([]string, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	if len(doc.Find(".no-result").Nodes) != 0 {
		return nil, ErrNotFound
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

func GetLoupanInfo(url string) (*LoupanInfo, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	sel := doc.Find(".module-loupan-circum")

	id, ok := sel.Attr("data-group_id")
	if !ok {
		return nil, ErrNotFound
	}
	name, ok := sel.Attr("data-house_name")
	if !ok {
		return nil, ErrNotFound
	}
	latStr, ok := sel.Attr("data-lat")
	if !ok {
		return nil, ErrNotFound
	}
	lonStr, ok := sel.Attr("data-lng")
	if !ok {
		return nil, ErrNotFound
	}
	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		return nil, err
	}
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		return nil, err
	}

	return &LoupanInfo{
		Id: id,
		Name: name,
		Url: url,
		Lat: lat,
		Lon: lon,
	}, nil
}
