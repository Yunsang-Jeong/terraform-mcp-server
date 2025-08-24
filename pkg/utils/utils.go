package utils

func IsInList(val string, allowed []string) bool {
	for _, v := range allowed {
		if v == val {
			return true
		}
	}
	return false
}
