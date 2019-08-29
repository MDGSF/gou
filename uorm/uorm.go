package uorm

// StringArray2SQLString 把数组转换为可以放到 mysql 的 where in () 中的字符串
func StringArray2SQLString(in []string) string {
	result := ""
	for k := range in {
		if k == 0 {
			result += "\"" + in[k] + "\""
		} else {
			result += ", \"" + in[k] + "\""
		}
	}
	return result
}
