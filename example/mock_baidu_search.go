// an example to mock search in baidu
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var data = `
{
	"load_images": false, 
	"url": "http://www.baidu.com", 
	"headers": {"User-Agent": ""}, 
	"timeout": 120, 
	"use_gzip": true, 
	"allow_redirects": true, 
	"method": "GET",
	"js_script":"function(){document.getElementById('kw').setAttribute('value', 'github');document.getElementById('su').click();}",
	"js_run_at":"document-end"
}`

func main() {
	buffer := bytes.NewBuffer([]byte(data))
	res, err := http.Post("http://localhost:2000", "application/json;charset=utf-8", buffer)
	if err != nil {
		panic(err)
	}

	byte_data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(byte_data))
	res.Body.Close()

	var result map[string]interface{}
	err = json.Unmarshal(byte_data, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(result["content"])
}
