package dragonvalid

import "unicode"

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
