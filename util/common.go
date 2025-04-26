package util

import (
	"fmt"
	"strings"
	"time"

	uuid "github.com/google/uuid"
	"github.com/gosimple/slug"
)

func FindInArray(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func StrPadLeft(input string, padLength int, padString string) string {
	output := padString

	for padLength > len(output) {
		output += output
	}

	if len(input) >= padLength {
		return input
	}

	return output[:padLength-len(input)] + input
}

func TrimmedDays() []string {
	return []string{"Mon,Tue,Wed,Thu,Fri,Sat,Sun"}
}

func NameOfDays() []string {
	return []string{"Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday"}
}

func ArrToStrDelimiter(arrs []string, delimiter string) string {
	arrCheck := []string{}
	result := ""
	if delimiter == "" {
		delimiter = ","
	}
	for _, arr := range arrs {
		if !FindInArray(arrCheck, arr) {
			arrCheck = append(arrCheck, arr)
			result += "'" + arr + "'" + delimiter
		}
	}
	return strings.TrimSuffix(result, delimiter)
}

func FirstSaturday(year int, month time.Month) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return (6 - int(t.Weekday())) + t.Day()
}

func InArray(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}

// Get UUID
func GetUUID(input string) uuid.UUID {
	id, err := uuid.Parse(input)
	if err != nil {
		return id
	}
	return id
}

// Get Slug
func GetSlug(ch string) string {
	text := slug.Make(ch)
	return fmt.Sprintf("%s-%d", text, time.Now().Unix())
}
