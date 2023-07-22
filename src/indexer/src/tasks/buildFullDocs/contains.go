package buildfulldocs

func contains(lst []string, str string) bool {
	for _, item := range lst {
		if item == str {
			return true
		}
	}
	return false
}
