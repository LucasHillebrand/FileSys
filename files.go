package filesys

import (
	defaultfunc "Creator/DefaultFunc"
	"os"
)

func GetFile(FileName string) string {
	out := ""
	data, err := os.ReadFile(FileName)
	if err == nil {
		for i := 0; i < len(data); i++ {
			out += string(data[i])
		}
	}
	return out
}

func GetCsvFile(FileName string) [][]string {
	lines := GetLines(FileName)
	out := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		out[i] = defaultfunc.Split(lines[i], ",")
	}
	return out
}

func GetLines(FileName string) []string {
	data := GetFile(FileName)
	out := defaultfunc.Split(data, "\n")
	return out
}
