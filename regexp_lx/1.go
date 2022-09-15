package main

import (
	"github.com/davecgh/go-spew/spew"
	"regexp"
)

type reg struct {
	re *regexp.Regexp
}

func main() {
	s := ":id(.*)"

	//b := regexp.MustCompile(`:id(.*)`)

	a := new(reg)

	a.re = regexp.MustCompile(`:id(.*)`)

	//
	spew.Dump(s)
	spew.Dump(a)
}
