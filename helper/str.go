package helper

import (
	"fmt"
	"regexp"
	"strings"
)

// Checks if a string contains, does not contain a specific string, and conforms to a regular expression.
// If the regular expression pattern is invalid, return true.
// If either condition is not met, return false.
func Filter(target string, include []string, exclude []string, regex string) bool {
	result := true
	if len(include) > 0 {
		re, err := regexp.Compile(fmt.Sprintf(".*(%s).*", strings.Join(include, "|")))
		if err != nil {
			return true
		}
		result = result && re.MatchString(target)
	}
	if len(exclude) > 0 {
		re, err := regexp.Compile(fmt.Sprintf(".*(%s).*", strings.Join(exclude, "|")))
		if err != nil {
			return true
		}
		result = result && !re.MatchString(target)
	}
	if len(regex) > 0 {
		re, err := regexp.Compile(regex)
		if err != nil {
			return true
		}
		result = result && re.MatchString(target)
	}
	return result
}
