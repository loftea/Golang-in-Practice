package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	XMLName xml.Name
	Status  Status
}

func main() {
	response, _ := http.Get("http://weibo.com/u/1686546714.xml")
	user := User{xml.Name{Space: "", Local: "user"}, Status{""}}

	body, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal([]byte(body), &user)

	fmt.Printf("status:%s", user.Status.Text)

}
