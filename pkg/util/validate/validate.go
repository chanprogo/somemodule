package validate

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// 字符串不能为空
func Required(value string) error {
	if value != "" {
		return nil
	}
	return errors.New(value + "不能为空")
}

// 最小值
func Min(value string, min int) error {
	if num, ok := strconv.Atoi(value); ok == nil {
		if num >= min {
			return nil
		}
	}
	return errors.New(value + "最小为" + strconv.Itoa(min))
}

// 最大值
func Max(value string, max int) error {
	if num, ok := strconv.Atoi(value); ok == nil {
		if num <= max {
			return nil
		}
	}
	return errors.New(value + "最大为" + strconv.Itoa(max))
}

// 数值范围
func Range(value string, min int, max int) error {
	if num, ok := strconv.Atoi(value); ok == nil {
		if num >= min && num <= max {
			return nil
		}
	}
	return errors.New(value + "范围为" + strconv.Itoa(min) + "-" + strconv.Itoa(max))
}

// 最小长度
func MinLength(value string, min int) error {
	if len([]rune(value)) >= min {
		return nil
	}
	return errors.New(value + "最小长度为" + strconv.Itoa(min))
}

// 最大长度
func MaxLength(value string, max int) error {
	if len([]rune(value)) <= max {
		return nil
	}
	return errors.New(value + "最大长度为" + strconv.Itoa(max))
}

// 长度范围
func Length(value string, length int) error {
	if len([]rune(value)) == length {
		return nil
	}
	return errors.New("长度应为" + strconv.Itoa(length))
}

// 字母或数字的组合
func AlphaNumeric(value string) error {
	for _, v := range value {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
			return errors.New(value + "只能为字母或数字")
		}
	}
	return nil
}

// 字母或数字的组合
func AlphaNumericHua(value string) error {
	for _, v := range value {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') && v != '_' {
			return errors.New(value + "只能为字母或数字或下划线")
		}
	}
	return nil
}

// 数字的组合
func Numeric(value string) error {
	for _, v := range value {
		if '9' < v || v < '0' {
			return errors.New(value + "只能为数字")
		}
	}
	return nil
}

// 字母或数字的组合
func AlphaNumericBlank(value string) error {
	for _, v := range value {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') && v != ' ' {
			return errors.New(value + "只能为字母或数字")
		}
	}
	return nil
}

// 字母或数字的组合
func AlphaNumericHan(value string) error {
	for _, v := range value {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
			if !unicode.Is(unicode.Scripts["Han"], v) {
				return errors.New(value + "只能为字母或数字或汉字")
			}
		}
	}

	return nil
}

// 指定字符串
func Enum(value string, enums string) error {
	// 拆分字符串
	each := strings.Split(enums, ",")
	for k, _ := range each {
		if value == strings.TrimSpace(each[k]) {
			return nil
		}
	}
	return errors.New(value + "不在指定范围内")
}

func VerifyEmailFormat(email string) error {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	status := reg.MatchString(email)
	if !status {
		return errors.New(email + "输入的邮箱格式不正确")
	}
	return nil
}

func VerifyMobileFormat(mobileNum string) error {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	status := reg.MatchString(mobileNum)
	if !status {
		return errors.New(mobileNum + "输入的手机格式不正确")
	}
	return nil
}

func VerifyAlipayFormat(value string) error {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	status1 := reg.MatchString(value)

	pattern = "^((13[0-9])|(14[5,7])|(15[0-9])|(17[0-9])|(18[0-9])|(16[0-9])|(19[0-9])|(147))\\d{8}$"
	reg = regexp.MustCompile(pattern)
	status2 := reg.MatchString(value)

	if !status1 && !status2 {
		return errors.New(value + "输入的帐号格式不正确")
	}
	return nil
}
