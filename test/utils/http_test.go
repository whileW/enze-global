package test

import (
	"github.com/whileW/enze-global/utils"
	"testing"
)

func TestStringResult(t *testing.T)  {
	resp,err := utils.HttpStringResult(utils.Post("http://www.baidu.com","",nil))
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(resp)
}

func TestPost(t *testing.T)  {
	resp,err := utils.Post("http://www.baidu.com","",nil)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(resp))
}

type test_req_body_struct struct {
	Test 		string			`json:"test"`
}
func TestPostWithJson(t *testing.T)  {
	test_req_body := test_req_body_struct{
		Test:"test",
	}
	resp,err := utils.PostWithJson("http://www.baidu.com",test_req_body)
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(resp))
}

func TestGet(t *testing.T)  {
	resp,err := utils.Get("http://www.baidu.com")
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(resp))
}

func TestHttpClient(t *testing.T)  {
	resp,err := utils.HttpStringResult(
		utils.NewHttpClient(
			utils.GET,
			"http://www.baidu.com",
			nil).
			Do())
	if err != nil {
		t.Log(err.Error())
	}
	t.Log(string(resp))
}