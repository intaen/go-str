package str

import "strings"

// CutStr is a function to get first or last data. Return result in string.
// Pos means position. -1 to get first data and 1 to get last data.
func CutStr(data string, pos, strlen int) string {
	var str string
	if pos == -1 {
		str = data[:strlen]
	} else if pos == 1 {
		str = data[strings.LastIndex(data, "")-strlen:]
	}

	return str
}
