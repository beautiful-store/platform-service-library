package library

func Contains(items []string, key string) bool {
	for _, item := range items {
		if item == key {
			return true
		}
	}
	return false
}
