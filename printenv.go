package main

import (
	"fmt"
	"os"
	"strings"
)

// GetEnvDict - return a dictionary of environment variables and int representing the maximum length of the keys
func GetEnvDict() (map[string]string, int) {
	return GetEnvDictMatch("")
}

func GetEnvDictMatch(searchterm string) (map[string]string, int) {
	ret := map[string]string{}
	maxlen := 0
	for _, val := range os.Environ() {
		pieces := strings.SplitN(val, "=", 2)
		if searchterm != "" && !strings.Contains(strings.ToLower(pieces[0]), strings.ToLower(searchterm)) {
			continue
		}
		ret[pieces[0]] = pieces[1]
		if len(pieces[0]) > maxlen {
			maxlen = len(pieces[0])
		}
	}
	return ret, maxlen
}

// StringToSize - reformat string s to size sz. If s is larger or equal to s, then return s. Otherwise,
// return a space padded string starting with s
func StringToSize(s string, sz int) string {
	if len(s) >= sz {
		return s
	}
	return fmt.Sprintf("%s%s", s, strings.Repeat(" ", sz-len(s)))
}

func ReplaceColon(s string, sz int) string {
	if len(s) > 0 && string(s[0]) == ":" {
		s = s[1:]
	}
	return strings.Replace(s, ":",
		fmt.Sprintf("\n%s", strings.Repeat(" ", sz)), -1)

}

func main() {
	args := os.Args
	searchterm := ""
	if len(args) >= 2 {
		searchterm = args[1]
	}
	envdict, sz := GetEnvDictMatch(searchterm)
	for key, val := range envdict {
		//sz2 := sz - len(key)

		val = ReplaceColon(val, sz+4)
		fmt.Println(StringToSize(key, sz), "= ", val)
		fmt.Println("")
	}
}
