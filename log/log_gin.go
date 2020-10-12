package log

type GinLog struct {}

func (l *GinLog)Write(p []byte) (n int, err error) {
	log.Info(string(p))
	return len(p),nil
}

type GinErrLog struct {}

func (l *GinErrLog)Write(p []byte) (n int, err error) {
	log.Error(string(p))
	return len(p),nil
}