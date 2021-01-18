package log

type GinErrLog struct {}
func (l *GinErrLog)Write(p []byte) (n int, err error) {
	GetLoger().Errorw("gin error log","msg",string(p))
	return len(p),nil
}

//禁用gin默认日志
type DisableGinDefaultLog struct {}
func (l *DisableGinDefaultLog)Write(p []byte) (n int, err error) {
	return len(p),nil
}


