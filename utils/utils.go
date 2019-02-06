package utils

import (
	"os"
	"strings"
)

// LoadProgress prints '#' every 100 chunks
func LoadProgress() func() string {
	i := 0
	j := 0
	return func() (s string) {
		if i%100 == 0 {
			s = "#"
			j++
			if j%70 == 0 {
				s += "\n"
			}
		}
		i++
		return s
	}
}

// CreateDirs recursively creates directories based on given path
func CreateDirs(path string) error {
	dirs := strings.Split(path, string(os.PathSeparator))
	dirs = dirs[:len(dirs)-1]
	err := os.MkdirAll(strings.Join(dirs, string(os.PathSeparator)), 0755)
	return err
}
