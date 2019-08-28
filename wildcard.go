package utils

// IsStringWildCardMatchArray 在数组 array 中查找字符串 s，用通配符匹配。
func IsStringWildCardMatchArray(s string, array []string) bool {
	for _, item := range array {
		if WildCardMatch(item, s) {
			return true
		}
	}
	return false
}

/*
WildCardMatch 通配符匹配字符串
?: 匹配 0 个或者 1 个字符
*: 匹配 0 个，1 个或者任意多个字符
+: 匹配 1 个字符或者任意个字符，也就是至少 1 个字符

WildCardMatch("a?c", "abc") == true
WildCardMatch("a?c", "ac") == true

WildCardMatch("a*c", "ac") == true
WildCardMatch("a*c", "abc") == true
WildCardMatch("a*c", "abbbc") == true

WildCardMatch("a+c", "ac") == false
WildCardMatch("a+c", "abc") == true
WildCardMatch("a+c", "abbbc") == true
*/
func WildCardMatch(pattern string, name string) bool {
	if len(pattern) == 0 {
		return pattern == name
	}
	if pattern == "*" {
		return true
	}
	rpattern := make([]rune, 0, len(pattern))
	rname := make([]rune, 0, len(name))
	for _, r := range pattern {
		rpattern = append(rpattern, r)
	}
	for _, r := range name {
		rname = append(rname, r)
	}
	return WildCardDeepMatchRune(rpattern, rname, false)
}

func WildCardDeepMatchRune(pattern, name []rune, plus bool) bool {
	for len(pattern) > 0 {
		switch pattern[0] {
		case '?':
			/*
			 * '?' 表示 0 个或者 1 个字符
			 */

			return WildCardDeepMatchRune(pattern[1:], name, plus) ||
				(len(name) > 0 && WildCardDeepMatchRune(pattern[1:], name[1:], plus))
		case '*':
			/*
			 * '*' 表示 0 个，1 个字符或者任意个字符
			 */

			return WildCardDeepMatchRune(pattern[1:], name, plus) ||
				(len(name) > 0 && WildCardDeepMatchRune(pattern, name[1:], plus))

		case '+':
			/*
			 * '+' 表示 1 个字符或者任意个字符，至少一个字符
			 *
			 * plus = true 表示上一次匹配的 pattern 字符是 '+' 并且 '+'
			 * 已经匹配了至少一个字符了。
			 */

			if plus {
				return WildCardDeepMatchRune(pattern[1:], name, false) ||
					(len(name) > 0 && WildCardDeepMatchRune(pattern, name[1:], true))
			} else {
				return len(name) > 0 && WildCardDeepMatchRune(pattern, name[1:], true)
			}

		default:
			if len(name) == 0 || pattern[0] != name[0] {
				return false
			}
		}

		pattern = pattern[1:]
		name = name[1:]
	}

	return len(name) == 0 && len(pattern) == 0
}
