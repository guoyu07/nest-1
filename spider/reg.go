package main

import (
	"os"
	"bufio"
	"regexp"
)

var (
	regs = make([]*regexp.Regexp, 0)
	aimReg = regexp.MustCompile(`.*ershoufang/sh[0-9]*\.html$`)
)

func init() {
	f, err := os.Open("regexp.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		r, err := regexp.Compile(line)
		if err != nil {
			panic(err)
		}
		regs = append(regs, r)
	}
}

func March(s string) bool {
	for _, v := range regs {
		if v.MatchString(s) {
			return true
		}
	}
	return false
}

func IsAim(s string) bool {
	return aimReg.MatchString(s)
}
