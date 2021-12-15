package util

import "strings"

func GetAbsString(s,prefix,suffix string) string{
	s=strings.TrimSuffix(strings.TrimPrefix(s,prefix),suffix)
	return s
}

