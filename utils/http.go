package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpStringResult(result []byte,err error) (string,error) {
	if err != nil {
		return "",err
	}
	return string(result),err
}

func Post(url, contentType string, req_body []byte) ([]byte,error) {
	if contentType == "" {
		 contentType = "application/json"
	}
	resp,err := http.Post(url,contentType,bytes.NewReader(req_body))
	defer resp.Body.Close()
	if err != nil {
		return nil,err
	}
	return ioutil.ReadAll(resp.Body)
}
func PostWithJson(url, contentType string, req_body interface{}) (result []byte,err error) {
	req_body_byte := []byte{}
	if req_body != nil {
		req_body_byte,err = json.Marshal(req_body)
		if err != nil {
			return nil,err
		}
	}else {
		req_body_byte = nil
	}
	return Post(url,contentType,req_body_byte)
}

func Get(url string) ([]byte,error) {
	resp,err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type http_client struct {
	client  		*http.Client
	req 			*http.Request
	content_type	string
}
type http_method string
var (
	GET 	http_method = "GET"
	POST	http_method = "POST"
	DELETE	http_method = "DELETE"
	PUT		http_method = "PUT"
)
func NewHttpClient(method http_method,url string, req_body []byte) *http_client {
	hc := &http_client{}

	hc.client = &http.Client{}

	hc.req, _ = http.NewRequest(string(method),	url,bytes.NewBuffer(req_body))
	return hc
}
func NewHttpClientWithJson(method http_method,url string, req_body interface{}) (*http_client,error) {
	var err error
	req_body_byte := []byte{}
	if req_body != nil {
		req_body_byte,err = json.Marshal(req_body)
		if err != nil {
			return nil,err
		}
	}else {
		req_body_byte = nil
	}
	return NewHttpClient(method,url,req_body_byte),nil
}
func (hc *http_client)DisableCompression() *http_client {
	tr := &http.Transport{
		DisableCompression: true,
	}
	hc.client.Transport = tr
	return hc
}
func (hc *http_client)SetHeader(k,v string) *http_client {
	lk := strings.ToLower(k)
	if lk == "content_type" || lk == "contenttype" || lk == "content-type" {
		hc.SetConetentType(v)
		return hc
	}
	hc.req.Header.Add(k,v)
	return hc
}
func (hc *http_client)DelHeader(k string) *http_client {
	hc.req.Header.Del(k)
	return hc
}
func (hc *http_client)SetConetentType(content_type string) *http_client {
	switch content_type {
	case "json":
		content_type = "application/json;"
		break
	case "form-data":
		content_type = "multipart/form-data;"
		break
	}
	hc.content_type = content_type
	return hc
}
func (hc *http_client)Do() ([]byte,error) {
	if hc.content_type == "" {
		hc.SetConetentType("json")
	}
	hc.SetHeader("Content-Type",hc.content_type)

	resp, err := hc.client.Do(hc.req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}