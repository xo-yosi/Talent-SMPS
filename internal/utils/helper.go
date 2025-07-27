package utils

func GenerateNextCoordinatorID(lastID int) int {
	var num int
	if lastID == 0{
		return 0000
	}
	num++
	return num
}
