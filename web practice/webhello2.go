package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func HelloServer(ww http.ResponseWriter, req *http.Request) {
	fmt.Println("In HelloServer")
	sl := strings.Split(req.URL.Path[1:], "/")
	if len(sl) >= 2 {
		if sl[0] == "hello" {
			fmt.Fprintf(ww, "Hello,"+sl[1])
		} else if sl[0] == "shouthello" {
			fmt.Fprintf(ww, "Hello,"+strings.ToUpper(sl[1]))
		}
	}
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
