package main

import (
	"io"
	"log"
)

func main(){
	//var e error
	//errors.Wrap(e, "read failed")
}

func Write(w io.Writer,buf []byte) error {
	_ , err := w.Write(buf)
	if err != nil {
		log.Println("test success",err)
		return err
	}
	return nil
}
