package ghelp

//IsNumeric  判断是否是数字
func IsNumeric(item interface{}) bool {
	switch item.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true

	case string:
		return false
	default:
		return false
	}
}
