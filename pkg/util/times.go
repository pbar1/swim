package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseSeconds(t string) (float64, error) {
	spl := strings.Split(t, ":")
	if len(spl) == 1 {
		sec, err := strconv.ParseFloat(spl[0], 64)
		if err != nil {
			return 0, err
		}
		return sec, nil
	} else if len(spl) == 2 {
		sec, err := strconv.ParseFloat(spl[1], 64)
		if err != nil {
			return 0, err
		}
		min, err := strconv.Atoi(spl[0])
		if err != nil {
			return 0, err
		}
		return float64(min)*60.0 + sec, nil
	}
	return 0, fmt.Errorf("unable to parse swim time string: %s", t)
}
