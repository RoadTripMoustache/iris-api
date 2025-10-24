package utils

// ContainsKey checks if a specific key exists in a map.
// This function iterates through the map keys to find a match.
// Parameters:
//   - m: The map to search in
//   - key: The key to search for
//
// Returns:
//   - bool: true if the key exists in the map, false otherwise
func ContainsKey(m map[string][]string, key string) bool {
	isPresent := false

	for k := range m {
		if key == k {
			isPresent = true
			break
		}
	}

	return isPresent
}
