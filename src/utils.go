package src

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Float642String(f float64) string {
	return strconv.FormatFloat(f,'f', 3, 64)
}


func PostRequest(url string, xmlParam string) (body []byte, err error) {
	resp, err := http.Post(url,
		"application/xml",
		strings.NewReader(xmlParam))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("Post Error:", err)
		return
	}

	//content = string(body)
	return
}