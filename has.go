package dragonvalid

import (
	"strings"
	"unicode"
)

// 必须包含字母
func (v Engine) hasLetter(customError ...string) Engine {
	return pushQueue(&v, func(v *Engine) *Engine {
		for _, rv := range v.value {
			if unicode.IsLower(rv) || unicode.IsUpper(rv) {
				return v
			}
		}

		v.err = setError(v, "必须包含字母", customError...)
		return v
	})
}

// 必须包含小写字母
func (v Engine) hasLower(customError ...string) Engine {
	return pushQueue(&v, func(v *Engine) *Engine {
		for _, rv := range v.value {
			if unicode.IsLower(rv) {
				return v
			}
		}

		v.err = setError(v, "必须包含小写字母", customError...)
		return v
	})
}

// 必须包含大写字母
func (v Engine) hasUpper(customError ...string) Engine {
	return pushQueue(&v, func(v *Engine) *Engine {
		for _, rv := range v.value {
			if unicode.IsUpper(rv) {
				return v
			}
		}

		v.err = setError(v, "必须包含大写字母", customError...)
		return v
	})
}

// TODO:必须包含数字
func (v Engine) HasNumber(customError ...string) Engine {
	return pushQueue(&v, func(v *Engine) *Engine {
		for _, rv := range v.value {
			if unicode.IsDigit(rv) {
				return v
			}
		}

		v.err = setError(v, "必须包含数字", customError...)
		return v
	})
}

// TODO:必须包含符号
func (v Engine) HasSymbol(customError ...string) Engine {
	return pushQueue(&v, func(v *Engine) *Engine {
		for _, rv := range v.value {
			if unicode.IsSymbol(rv) {
				return v
			}
		}

		v.err = setError(v, "必须包含符号", customError...)
		return v
	})
}

// TODO:必须包含特定的字符串
func (v Engine) HasString(sub string, customError ...string) Engine {
	return pushQueue(&v, func(v *Engine) *Engine {
		if !strings.Contains(v.value, sub) {
			v.err = setError(v, "必须包含特定的字符串", customError...)
			return v
		}
		return v
	})
}

// TODO:必须包含指定前缀

// TODO:必须包含指定后缀

//TODO:6-20 位字符

// TODO: 必须包含大小英文字母+符号并且长度为 6-20 位

// TODO: 正则匹配
