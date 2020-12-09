package log

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"time"
)

const MAX_PRINT_GIN_RESP_BODY_LEN  =  512

type GinBodyLoger struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
	RequestBody	[]byte
}

func (w GinBodyLoger) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

//启用gin 请求日志
func EnableGinLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		kvs := []interface{}{}
		start := time.Now()

		gin_body_loger := GinBodyLoger{
			bodyBuf: bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		//req body
		if d,err := ioutil.ReadAll(c.Request.Body);err == nil{
			gin_body_loger.RequestBody = d
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(d))
		}

		//resp body
		c.Writer = gin_body_loger

		//开关
		c.Set("log_resp_body_writer",gin_body_loger)
		c.Set("is_log_resp_body",true)
		c.Set("is_log_req_body",true)
		//path := c.Request.URL.Path
		//raw := c.Request.URL.RawQuery

		c.Next()

		//req infos
		ip := c.ClientIP()
		method := c.Request.Method
		path := c.Request.RequestURI
		kvs = append(kvs,"path",path,"ip",ip,"method",method)

		//处理时间
		over := time.Now()
		latency := over.Sub(start)
		kvs = append(kvs, "req_time",start,"latency",latency)

		//req body
		if is_log := c.GetBool("is_log_req_body");is_log {
			req_body := string(gin_body_loger.RequestBody)
			kvs = append(kvs, "req_body",req_body)
		}
		
		//resp body
		if is_log := c.GetBool("is_log_resp_body");is_log {
			str_body := strings.Trim(gin_body_loger.bodyBuf.String(), "\n")
			if len(str_body) > MAX_PRINT_GIN_RESP_BODY_LEN {
				str_body = str_body[:MAX_PRINT_GIN_RESP_BODY_LEN-1]
			}
			kvs = append(kvs, "resp_body",str_body)
		}
		status_code := c.Writer.Status()
		kvs = append(kvs,"status_code",status_code)
		log.Infow("请求日志",kvs...)
	}
}
//禁用gin resp body 日志
func DisableGinRespBodyLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_log_resp_body",false)
		if v,exist:=c.Get("log_resp_body_writer");exist{
			if log,ok := v.(GinBodyLoger);ok {
				c.Writer = log.ResponseWriter
			}
		}
		c.Next()
	}
}
//禁用gin req body 日志
func DisableGinReqBodyLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("is_log_req_body",false)
		c.Next()
	}
}


type GinErrLog struct {}
func (l *GinErrLog)Write(p []byte) (n int, err error) {
	log.Errorw("gin error log","msg",string(p))
	return len(p),nil
}

//禁用gin默认日志
type DisableGinDefaultLog struct {}
func (l *DisableGinDefaultLog)Write(p []byte) (n int, err error) {
	return len(p),nil
}


