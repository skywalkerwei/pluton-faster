package ghelp

import (
	"strconv"
)

//Tostring 转成字符串 多个字符串可拼接
func ToString(args ...interface{}) string {
	result := ""
	for _, arg := range args {
		switch val := arg.(type) {
		case string:
			result += val
		case uint:
			result += strconv.Itoa(int(val))
		case int8:
			result += strconv.Itoa(int(val))
		case int16:
			result += strconv.Itoa(int(val))
		case int32:
			result += strconv.Itoa(int(val))
		case int:
			result += strconv.Itoa(val)
		case int64:
			result += strconv.Itoa(int(val))
		case float32:
			result += strconv.FormatFloat(float64(val), 'f', -1, 32)
		case float64:
			result += strconv.FormatFloat(val, 'f', -1, 64)
		case bool:
			result += strconv.FormatBool(val)
		}
	}
	return result
}

//ToInt 转成数字类型 以及浮点四舍五入取整
func ToInt(arg interface{}) int64 {
	var result int64 = 0
	switch val := arg.(type) {
	case string:
		result, err := strconv.Atoi(val)
		if err != nil {
			result = 0
		}
		return int64(result)
	case float32:
		result, err := strconv.Atoi(strconv.FormatFloat(float64(val), 'f', 0, 32))
		if err != nil {
			result = 0
		}
		return int64(result)
	case float64:
		result, err := strconv.Atoi(strconv.FormatFloat(val, 'f', 0, 64))
		if err != nil {
			result = 0
		}
		return int64(result)
	case bool:
		switch val {
		case true:
			return 1
		case false:
			return 0
		}
	case int8:
		return int64(val)
	case int16:
		return int64(val)
	case int32:
		return int64(val)
	case int64:
		return val
	}

	return result
}
