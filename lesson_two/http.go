package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)
	http.HandleFunc("/body/once", readBodyOnce)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return
	}
	// 类型转换，将 []byte 转换为 string
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

func form(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Fprintf(w, "query is %v\n", values)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi, this is home page")
}
