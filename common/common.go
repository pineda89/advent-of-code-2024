package common

import "os"

func GetInput(filepath string) string {
	cnt, _ := os.ReadFile(filepath)
	return string(cnt)
}
