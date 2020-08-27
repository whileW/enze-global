package utils

//切片操作

func SliceStringContains(a []string,b string) bool {
	for _,t := range a {
		if t == b {
			return true
		}
	}
	return false
}