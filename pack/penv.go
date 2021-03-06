package penv

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const Usage string = `Print out the environment to std.out`

// GetEnvDict - return a dictionary of environment variables and int representing the maximum length of the keys
func GetEnvDict(matchFunc MatchFunc) (map[string]string, int) {
	return GetEnvDictMatch("", matchFunc)
}

type MatchFunc func(rhs, lhs string) bool

func ExactMatch(rhs, lhs string) bool {
	if rhs == lhs {
		return true
	}
	return false
}

// does the lhs
func ContainsMatch(rhs, lhs string) bool {
	ct := strings.Contains(rhs, lhs)
	return ct
}

func GetEnvDictMatch(searchterm string, matchFunc MatchFunc) (map[string]string, int) {
	ret := map[string]string{}
	maxlen := 0
	for _, val := range os.Environ() {
		pieces := strings.SplitN(val, "=", 2)

		if searchterm != "" && !matchFunc(strings.ToLower(pieces[0]), strings.ToLower(searchterm)) {
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

// ValModeFunc type - function to modify the the value of the first input, given a
// second input which is the size of the size of the key.
type ValModFunc func(string, int) string

// PrintLineFunc - given a key, a size integer, and a value, print the string as you see fit.
type PrintLineFunc func(key string, sz int, val string)

// ReplaceColon - a function of tpe ValModFunc
func ReplaceColon(s string, sz int) string {
	if len(s) > 0 && string(s[0]) == ":" {
		s = s[1:]
	}
	return strings.Replace(s, ":",
		fmt.Sprintf("\n%s", strings.Repeat(" ", sz)), -1)

}

// PassThroughVMF - a ValModFunc which doesn't modify the input in any way.
func PassThroughVMF(s string, sz int) string { return s }

// FormatPrint - a function of type PrintLineFunc
func FormatPrint(key string, sz int, val string) {

	fmt.Println(StringToSize(key, sz), "= ", val)
}

func FormatPrintWithSep(key string, sz int, val string) {

	fmt.Println(StringToSize(key, sz), "= ", val)
	fmt.Println("")
}

// StrPair is a pair of strings with a First and Second item.
type StrPair struct {
	First  string
	Second string
}

type StrPairs []StrPair

func (slice StrPairs) Len() int {
	return len(slice)
}

func (slice StrPairs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice StrPairs) Less(i, j int) bool {
	return slice[i].First < slice[j].First
}

// PrintEnv - used to print the environment
func PrintEnv(searchterm string, matchFunc MatchFunc, valFunc ValModFunc, printLineFunc PrintLineFunc, doSort bool) {
	envdict, sz := GetEnvDictMatch(searchterm, matchFunc)
	var strpairs StrPairs
	if doSort {
		strpairs = make(StrPairs, len(envdict))
		cnt := 0
		for key, val := range envdict {
			val = valFunc(val, sz+4)
			strpairs[cnt] = StrPair{key, val}
			cnt++
		}
		sort.Sort(strpairs)
		for _, v := range strpairs {
			printLineFunc(v.First, sz, v.Second)
		}
	} else {

		for key, val := range envdict {
			val = valFunc(val, sz+4)
			printLineFunc(key, sz, val)
		}
	}
}
