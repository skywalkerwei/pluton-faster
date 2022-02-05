package ghelp

import (
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

func VerifyChineseName(name string) bool {
	expr := "^[\\p{Han}]{2,4}$"
	r, _ := regexp.Compile(expr)
	return r.MatchString(name)

}

func VerifyTelNum(name string) bool {
	expr := "^1\\d{10}$"
	r, _ := regexp.Compile(expr)
	return r.MatchString(name)

}

func VerifyIDCard(idCard string) bool {
	expr := "(^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}(\\d|x|X)$)|(^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$)"

	ret1, _ := verifyIDCardLastCode(idCard)
	if !ret1 {
		return false
	}

	r, _ := regexp.Compile(expr)
	return r.MatchString(idCard)
}

func verifyIDCardLastCode(id string) (bool, error) {
	lastOne, err := getCheckCode(id)
	if err != nil {
		return false, err
	}

	if string(id[17]) == lastOne {
		return true, nil
	}
	return false, nil
}

func getCheckCode(id string) (string, error) {
	var weight = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var code = []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	if len(id) != 18 {
		return "", fmt.Errorf("长度不对")
	}

	sum := 0
	for i := 0; i < 17; i++ {
		if unicode.IsDigit(rune(id[i])) {
			n, err := strconv.Atoi(string(id[i]))
			if err != nil {
				return "", err
			}
			sum += n * weight[i]
		} else {
			return "", fmt.Errorf("不是数字")
		}
	}

	m := sum % 11
	return code[m], nil
}
