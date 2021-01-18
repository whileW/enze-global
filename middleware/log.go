package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/whileW/enze-global/log"
	"io/ioutil"
	"strings"
	"time"
)

//最大打印gin resp_body长度
const MAX_PRINT_GIN_RESP_BODY_LEN  =  512

type ginRespBodyLoger struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}
func (w ginRespBodyLoger) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

//是否记录body日志
func is_log_req_body(c *gin.Context) bool {
	content_type := c.Request.Header.Get("Content-Type")
	if content_type != "" {
		content_type = strings.TrimSpace(strings.Split(content_type,";")[0])
		if strings.ToLower(c.Request.Method) == "post" && content_type == "application/json" {
			c.Set("is_log_req_body",true)
			return true
		}
	}
	c.Set("is_log_req_body",false)
	return false
}
func is_log_resp_body(c *gin.Context) bool {
	content_type := c.Writer.Header().Get("Content-Type")
	if content_type != "" {
		content_type = strings.TrimSpace(strings.Split(content_type, ";")[0])
		if content_type == "application/json" {
			c.Set("is_log_resp_body", true)
			return true
		}
	}
	c.Set("is_log_resp_body",false)
	return false
}

//gin请求日志中间件
func EnableGinLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		kvs := []interface{}{}
		start := time.Now()

		gin_body_loger := ginRespBodyLoger{
			bodyBuf: bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		//req body
		if is_log_req_body(c) {
			if d,err := ioutil.ReadAll(c.Request.Body);err == nil{
				kvs = append(kvs, "req_body",string(d))
				c.Request.Body = ioutil.NopCloser(bytes.NewReader(d))
			}
		}

		//resp body
		if is_log_resp_body(c) {
			c.Writer = gin_body_loger
			//c.Set("log_resp_body_writer",gin_body_loger)
		}

		c.Next()

		//req infos
		ip := c.ClientIP()
		method := c.Request.Method
		path := c.Request.RequestURI
		kvs = append(kvs,"path",path,"ip",ip,"method",method)

		//resp body
		if is_log_resp_body(c) {
			resp_body := strings.Trim(gin_body_loger.bodyBuf.String(), "\n")
			if len(resp_body) > MAX_PRINT_GIN_RESP_BODY_LEN {
				resp_body = resp_body[:MAX_PRINT_GIN_RESP_BODY_LEN-1]
			}
			kvs = append(kvs, "resp_body", resp_body)
		}
		status_code := c.Writer.Status()
		kvs = append(kvs,"status_code",status_code)

		//处理时间
		kvs = append(kvs, "req_time",start)
		over := time.Now()
		latency := over.Sub(start)
		log.GetLoger().WithDuration(latency).Infow("请求日志",kvs...)
	}
}