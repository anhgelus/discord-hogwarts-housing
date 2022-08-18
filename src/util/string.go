package util

import (
	"github.com/anhgelus/discord-hogwarts-housing/src/config"
	"strings"
)

func GetNumberOfChar(s string) int {
	list := []string{""}
	for i := 0; i < len(s); i++ {
		for _, j := range list {
			if j == s[i:i+1] {
				continue
			}
		}
		list = append(list, s[i:i+1])
	}
	return len(list)
}

func StartWith(m string, c string) bool {
	return strings.HasPrefix(m, config.Prefix+c)
}
