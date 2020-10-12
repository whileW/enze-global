package log

type GinLog struct {}

func (l *GinLog)Write(p []byte) (n int, err error) {
	log.Info(p)
	return len(p),nil
}

type GinErrLog struct {}

func (l *GinErrLog)Write(p []byte) (n int, err error) {
	log.Error(p)
	return len(p),nil
}