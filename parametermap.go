package netsuiteodbc

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

var _ fmt.Stringer = (*parameterMap)(nil)

type parameterMap map[string]string

func (c parameterMap) String() string {
	str := ""
	for key, value := range c {
		str = fmt.Sprintf("%s%s=%s;", str, key, value)
	}
	return str
}

func connStringToParameterMap(connStr string) (parameterMap, error) {
	m := make(parameterMap)
	r := strings.NewReader(connStr)

	readUntil := func(delimiter rune) (key string, eof bool) {
		buffer := bytes.NewBuffer(nil)
		parens := 0
		for {
			ch, _, err := r.ReadRune()
			if err == io.EOF {
				return buffer.String(), true
			}
			if ch == delimiter && parens == 0 {
				return buffer.String(), false
			} else if ch == '(' {
				parens++
			} else if ch == ')' {
				parens--
			}

			buffer.WriteRune(ch)
		}
	}

	for {
		key, eof := readUntil('=')
		if eof {
			break
		}
		value, eof := readUntil(';')
		m[key] = value
		if eof {
			break
		}
	}
	return m, nil
}
