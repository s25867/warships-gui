package gui

import (
	"strconv"
	"strings"
)

func rgbFromString(s string) (int, int, int, error) {
	var colors [3]int
	for i, e := range strings.Split(s, ",") {
		n, err := strconv.Atoi(e)
		if err != nil {
			return -1, -1, -1, err
		}
		colors[i] = n
	}
	return colors[0], colors[1], colors[2], nil
}
