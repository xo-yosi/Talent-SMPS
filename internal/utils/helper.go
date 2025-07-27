package utils

func GenerateNextCoordinatorID(lastID int) int {
	if lastID == 0 {
		return 100
	}
	lastID++
	return lastID
}
