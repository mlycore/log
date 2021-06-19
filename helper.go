package log

import "strings"

func getShortFileName(file string) string {
	index := strings.LastIndex(file, "/")
	return file[index+1:]
}
