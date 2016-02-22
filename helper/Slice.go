package helper

//Contains returns true if the needle is within the haystack.
func Contains(haystack []string, needle string) (bool, int) {
	for i, v := range haystack {
		if v == needle {
			return true, i
		}
	}
	return false, -1
}
