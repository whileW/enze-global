package utils

func IF(condition bool,trueval,falseval interface{}) interface{} {
	if condition {
		return trueval
	}
	return falseval
}
