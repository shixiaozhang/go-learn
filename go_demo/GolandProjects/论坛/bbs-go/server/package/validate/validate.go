package validate

import (
	"errors"
	"github.com/mlogclub/simple"
	"regexp"
	"strings"
)

// IsUsername 验证用户名合法性，用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头。
func IsUsername(username string) error {
	if simple.IsBlank(username) {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

// IsEmail 验证是否是合法的邮箱
func IsEmail(email string) (err error) {
	if simple.IsBlank(email) {
		err = errors.New("邮箱格式不符合规范")
		return
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return
}

// IsPassword 是否是合法的密码
func IsPassword(password, rePassword string) error {
	if simple.IsBlank(password) {
		return errors.New("请输入密码")
	}
	if simple.RuneLen(password) < 6 {
		return errors.New("密码过于简单")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}
	return nil
}

// IsURL 是否是合法的URL
func IsURL(url string) error {
	if simple.IsBlank(url) {
		return errors.New("URL格式错误")
	}
	indexOfHttp := strings.Index(url, "http://")
	indexOfHttps := strings.Index(url, "https://")
	if indexOfHttp == 0 || indexOfHttps == 0 {
		return nil
	}
	return errors.New("URL格式错误")
}
