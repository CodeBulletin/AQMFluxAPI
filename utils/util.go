package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func ExtractFirstNumber(input []byte) (int32, int, error) {
	str := string(input)

	re := regexp.MustCompile(`\d+`)
	match := re.FindString(str)

	if match == "" {
		return 0, -1, fmt.Errorf("no numbers found in the string")
	}

	number, err := strconv.Atoi(match)
	if err != nil {
		return 0, -1, fmt.Errorf("error converting string to integer: %v", err)
	}

	lastIndex := len(match) - 1

	return int32(number), int(lastIndex), nil
}

func ParseTime(ti string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, ti)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing time: %v", err)
	}
	return t, nil
}

func JoinStrings(strs []string, separator string) string {
	return strings.Join(strs, separator)
}

func SplitString(str string, separator rune) []string {
	return strings.Split(str, string(separator))
}