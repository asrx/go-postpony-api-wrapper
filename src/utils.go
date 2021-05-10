package src

import (
	"crypto/tls"
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
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Post(url,
		"application/xml",
		strings.NewReader(xmlParam))

	//resp, err := http.Post(url,
	//	"application/xml",
	//	strings.NewReader(xmlParam))

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