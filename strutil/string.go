// Use of this source code is governed by MIT license

// Package strutil implements some functions to manipulate string.
package strutil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelCase coverts string to camelCase string. Non letters and numbers will be ignored.
// Play: https://go.dev/play/p/9eXP3tn2tUy
func CamelCase(s string) string {
	var builder strings.Builder

	strs := splitIntoStrings(s, false)
	for i, str := range strs {
		if i == 0 {
			builder.WriteString(strings.ToLower(str))
		} else {
			builder.WriteString(Capitalize(str))
		}
	}

	return builder.String()
}

// Capitalize converts the first character of a string to upper case and the remaining to lower case.
// Play: https://go.dev/play/p/2OAjgbmAqHZ
func Capitalize(s string) string {
	result := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			result[i] = unicode.ToUpper(v)
		} else {
			result[i] = unicode.ToLower(v)
		}
	}

	return string(result)
}

// UpperFirst converts the first character of string to upper case.
// Play: https://go.dev/play/p/sBbBxRbs8MM
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)

	return string(r) + s[size:]
}

// LowerFirst converts the first character of string to lower case.
// Play: https://go.dev/play/p/CbzAyZmtJwL
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)

	return string(r) + s[size:]
}

// PadEnd pads string on the right side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/9xP8rN0vz--
func PadEnd(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return source + fill[0:size-len1]
}

// PadStart pads string on the left side if it's shorter than size.
// Padding characters are truncated if they exceed size.
// Play: https://go.dev/play/p/xpTfzArDfvT
func PadStart(source string, size int, padStr string) string {
	len1 := len(source)
	len2 := len(padStr)

	if len1 >= size {
		return source
	}

	fill := ""
	if len2 >= size-len1 {
		fill = padStr[0 : size-len1]
	} else {
		fill = strings.Repeat(padStr, size-len1)
	}
	return fill[0:size-len1] + source
}

// KebabCase coverts string to kebab-case, non letters and numbers will be ignored.
// Play: https://go.dev/play/p/dcZM9Oahw-Y
func KebabCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "-")
}

// UpperKebabCase coverts string to upper KEBAB-CASE, non letters and numbers will be ignored
// Play: https://go.dev/play/p/zDyKNneyQXk
func UpperKebabCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "-")
}

// SnakeCase coverts string to snake_case, non letters and numbers will be ignored
// Play: https://go.dev/play/p/tgzQG11qBuN
func SnakeCase(s string) string {
	result := splitIntoStrings(s, false)
	return strings.Join(result, "_")
}

// UpperSnakeCase coverts string to upper SNAKE_CASE, non letters and numbers will be ignored
// Play: https://go.dev/play/p/4COPHpnLx38
func UpperSnakeCase(s string) string {
	result := splitIntoStrings(s, true)
	return strings.Join(result, "_")
}

// Before returns the substring of the source string up to the first occurrence of the specified string.
// Play: https://go.dev/play/p/JAWTZDS4F5w
func Before(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[0:i]
}

// BeforeLast returns the substring of the source string up to the last occurrence of the specified string.
// Play: https://go.dev/play/p/pJfXXAoG_Te
func BeforeLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[0:i]
}

// After returns the substring after the first occurrence of a specified string in the source string.
// Play: https://go.dev/play/p/RbCOQqCDA7m
func After(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.Index(s, char)
	return s[i+len(char):]
}

// AfterLast returns the substring after the last occurrence of a specified string in the source string.
// Play: https://go.dev/play/p/1TegARrb8Yn
func AfterLast(s, char string) string {
	if s == "" || char == "" {
		return s
	}
	i := strings.LastIndex(s, char)
	return s[i+len(char):]
}

// IsString check if the value data type is string or not.
// Play: https://go.dev/play/p/IOgq7oF9ERm
func IsString(v any) bool {
	if v == nil {
		return false
	}
	switch v.(type) {
	case string:
		return true
	default:
		return false
	}
}

// Reverse returns string whose char order is reversed to the given string.
// Play: https://go.dev/play/p/adfwalJiecD
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Wrap a string with given string.
// Play: https://go.dev/play/p/KoZOlZDDt9y
func Wrap(str string, wrapWith string) string {
	if str == "" || wrapWith == "" {
		return str
	}
	var sb strings.Builder
	sb.WriteString(wrapWith)
	sb.WriteString(str)
	sb.WriteString(wrapWith)

	return sb.String()
}

// Unwrap a given string from anther string. will change source string.
// Play: https://go.dev/play/p/Ec2q4BzCpG-
func Unwrap(str string, wrapToken string) string {
	if str == "" || wrapToken == "" {
		return str
	}

	firstIndex := strings.Index(str, wrapToken)
	lastIndex := strings.LastIndex(str, wrapToken)

	if firstIndex == 0 && lastIndex > 0 && lastIndex <= len(str)-1 {
		if len(wrapToken) <= lastIndex {
			str = str[len(wrapToken):lastIndex]
		}
	}

	return str
}

// SplitEx split a given string which can control the result slice contains empty string or not.
// Play: https://go.dev/play/p/Us-ySSbWh-3
func SplitEx(s, sep string, removeEmptyString bool) []string {
	if sep == "" {
		return []string{}
	}

	n := strings.Count(s, sep) + 1
	a := make([]string, n)
	n--
	i := 0
	sepSave := 0
	ignore := false

	for i < n {
		m := strings.Index(s, sep)
		if m < 0 {
			break
		}
		ignore = false
		if removeEmptyString {
			if s[:m+sepSave] == "" {
				ignore = true
			}
		}
		if !ignore {
			a[i] = s[:m+sepSave]
			s = s[m+len(sep):]
			i++
		} else {
			s = s[m+len(sep):]
		}
	}

	var ret []string
	if removeEmptyString {
		if s != "" {
			a[i] = s
			ret = a[:i+1]
		} else {
			ret = a[:i]
		}
	} else {
		a[i] = s
		ret = a[:i+1]
	}

	return ret
}

// Substring returns a substring of the specified length starting at the specified offset position.
// Play: Todo
func Substring(s string, offset int, length uint) string {
	rs := []rune(s)
	size := len(rs)

	if offset < 0 {
		offset = size + offset
		if offset < 0 {
			offset = 0
		}
	}
	if offset > size {
		return ""
	}

	if length > uint(size)-uint(offset) {
		length = uint(size - offset)
	}

	str := string(rs[offset : offset+int(length)])

	return strings.Replace(str, "\x00", "", -1)
}
