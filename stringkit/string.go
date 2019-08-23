package stringkit

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// RemoveAllSpace -
func RemoveAllSpace(v string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, v)
}

// GBKToUTF8 gbk convert to utf8
func GBKToUTF8(r io.Reader) io.Reader {
	return transform.NewReader(r, simplifiedchinese.GBK.NewDecoder())
}

// https://siongui.github.io/2018/10/27/auto-detect-and-convert-html-encoding-to-utf8-in-go/
func DetermineEncodingFromReader(r io.Reader) (e encoding.Encoding, name string, certain bool, err error) {
	b, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return
	}

	e, name, certain = charset.DetermineEncoding(b, "")
	return
}

func StringNamedFormat(s string, m map[string]string) string {
	for k, v := range m {
		s = strings.Replace(s, fmt.Sprintf("{%s}", k), v, -1)
	}
	return s
}

// 按utf8字符数裁剪字符串
// https://stackoverflow.com/questions/46415894/golang-truncate-strings-with-special-characters-without-corrupting-data/46416000
func TruncateUTF8String(s string, limit int) string {
	result := s
	chars := 0
	for i := range s {
		if chars >= limit {
			result = s[:i]
			break
		}
		chars++
	}
	return result
}