package util

import (
	"strconv"
	"strings"
)

func JoinString(a []string, sep string) string {
	return strings.Join(a, sep)
}

func JoinInt(a []int, sep string) string {
	if len(a) == 0 {
		return ""
	}
	s := make([]string, len(a))
	for i := range a {
		s[i] = strconv.Itoa(a[i])
	}
	return strings.Join(s, sep)
}

func JoinInt64(a []int64, sep string) string {
	if len(a) == 0 {
		return ""
	}
	s := make([]string, len(a))
	for i := range a {
		s[i] = strconv.FormatInt(a[i], 10)
	}
	return strings.Join(s, sep)
}

func JoinInt32(a []int32, sep string) string {
	if len(a) == 0 {
		return ""
	}
	s := make([]string, len(a))
	for i := range a {
		s[i] = strconv.Itoa(int(a[i]))
	}
	return strings.Join(s, sep)
}
