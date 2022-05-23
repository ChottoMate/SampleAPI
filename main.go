package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Response struct {
	Staus   int      `json:"title"`
	Message string   `json:"message"`
	Results []Result `json:"results"`
}

type Result struct {
	Zipcode  string `json:"zipcode"`
	Prefcode string `json:"prefcode"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	Kana1    string `json:"kana1"`
	Kana2    string `json:"kana2"`
	Kana3    string `json:"kana3"`
}

func main() {
	values := url.Values{}
	values.Add("zipcode", "3350023")
	res, err := http.PostForm("https://zipcloud.ibsnet.co.jp/api/search", values)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	//var data []Response

	//if err := json.Unmarshal(body, &data); err != nil {
	//	log.Fatal(err)
	//}
	//for _, v := range data {
	//	fmt.Printf("%s %s %s\n", v.Staus, v.Message, v.Results[0])
	//}
}
