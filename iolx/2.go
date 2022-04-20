package main

import "io"

import (
	"fmt"
	"strings"
)

func main(){
	comment := "Package io provides basic interfaces to I/O primitives. " +
		"Its primary job is to wrap existing implementations of such primitives, " +
		"such as those in package os, " +
		"into shared public interfaces that abstract the functionality, " +
		"plus some other related primitives."
	reader1 := strings.NewReader(comment)
	buf1 := make([]byte, 7)
	n, err := reader1.Read(buf1)
	reader1.Reset(comment)
	num1 := int64(7)
	fmt.Printf("New a limited reader with reader1 and number %d ...\n", num1)
	reader2 := io.LimitReader(reader1, 7)
	buf2 := make([]byte, 10)
	for i := 0; i < 3; i++ {
		n, err = reader2.Read(buf2)
		executeIfNoErr(err, func() {
			fmt.Printf("Read(%d): %q\n", n, buf2[:n])
		})
	}
	fmt.Println()


}
