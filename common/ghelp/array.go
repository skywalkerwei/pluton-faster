package ghelp

// InArray 判断元素是否在数组中
// 支持string,int,int8,int16,int32,int64,float32,float64,bool类型
func InArray(item interface{}, items interface{}) bool {
	switch itemsValue := items.(type) {
	case []string:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int8:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int16:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int32:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []int64:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []float32:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []float64:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	case []bool:
		for _, v := range itemsValue {
			if v == item {
				return true
			}
		}
	}
	return false
}

// GenTree 将数组map生成tree
// items 将要排序的数组
// addEmptyChild child为空时是否返回空数组，默认true
// option[0] id别名，默认id
// option[1] pid别名，默认pid
// option[2] child别名，默认返回child
// example:
//  items := []map[string]interface{}{
//	   {"id": "1", "name": "1"},
//	   {"id": "2", "pid": "1", "name": "2"},
//	   {"id": "3", "pid": "1", "name": "3"},
//	   {"id": "4", "pid": "2", "name": "3"},
//   }
// JsonEncode(GenTree(items, true))
func GenTree(items []map[string]interface{}, addEmptyChild bool, option ...string) []map[string]interface{} {
	// 赋默认值
	_option := []string{"id", "pid", "child"}
	if len(option) > 0 {
		for k, v := range option {
			_option[k] = v
		}
	}

	newData := make(map[interface{}]map[string]interface{})
	for k, v := range items {
		if addEmptyChild {
			items[k][_option[2]] = []interface{}{}
		}
		newData[v[_option[0]]] = v
	}

	var tree []map[string]interface{}
	for _, v := range newData {
		_, find := v[_option[1]]
		var b bool
		if find {
			if _, ok := newData[v[_option[1]]]; ok {
				b = true
			}
		}
		if b {
			newData[v[_option[1]]][_option[2]] = append(newData[v[_option[1]]][_option[2]].([]interface{}), newData[v[_option[0]]])
		} else {
			tree = append(tree, newData[v[_option[0]]])
		}
	}
	return tree
}

//获取map里面所有键名
func ArrayKeys(items map[interface{}]interface{}) interface{} {
	i, keys := 0, make([]interface{}, len(items))
	for key := range items {
		keys[i] = key
		i++
	}
	return keys
}
