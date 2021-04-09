package utils

func FindSliceByKey(slice []string, index int, defaultValue string) string {
	for i, item := range slice {
		if i == index {
			return item
		}
	}

	return defaultValue
}
