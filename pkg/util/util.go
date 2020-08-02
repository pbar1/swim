package util

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var reg *regexp.Regexp

func init() {
	var err error
	reg, err = regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
}

func ParseSeconds(t string) (float64, error) {
	spl := strings.Split(t, ":")
	if len(spl) == 1 {
		sec, err := strconv.ParseFloat(spl[0], 64)
		if err != nil {
			return 0, fmt.Errorf(`parsing time string "%s": %v`, t, err)
		}
		return sec, nil
	} else if len(spl) == 2 {
		sec, err := strconv.ParseFloat(spl[1], 64)
		if err != nil {
			return 0, fmt.Errorf(`parsing time string "%s": %v`, t, err)
		}
		min, err := strconv.Atoi(spl[0])
		if err != nil {
			return 0, fmt.Errorf(`parsing time string "%s": %v`, t, err)
		}
		return float64(min)*60.0 + sec, nil
	}
	return 0, fmt.Errorf("unable to parse swim time string: %s", t)
}

func RemoveSpecialChars(s string) string {
	return reg.ReplaceAllString(s, "")
}
