package main

import (
	"io"
	"net/http"
)

const form = `
<html><body>
<form action="#" method="post" name="bar">
	<input type="text" name="in" />
	<input type="submit" value="submit"/>
</form>
</body></html>
`

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	//点击提交按钮
	case "POST":
		io.WriteString(w, r.FormValue("in"))
		//浏览器请求
	case "GET":
		io.WriteString(w, form)
	}
}

func main() {
	http.HandleFunc("/hello", SimpleServer)
	http.HandleFunc("/test", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
