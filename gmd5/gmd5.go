package gmd5

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/cnlesscode/gotool/random"
)

// 将字符串转换为自定义 md5 密码
func ToMd5Pwd(str string) string {
	str = Md5(Md5(str))
	offset := random.RangeIntRand(11, 30)
	randNum := random.RangeIntRand(10, 99)
	strF := str[0:offset]
	strB := str[offset:]
	newpass := strF + strconv.Itoa(int(randNum)) + strB + strconv.Itoa(int(offset))
	return newpass
}

// 将密码翻译为2次 md5
func PwdToMd5(pwd string) string {
	offsetStr := pwd[len(pwd)-2:]
	offsetNum, _ := strconv.Atoi(offsetStr)
	pwd = pwd[0 : len(pwd)-2]
	strF := pwd[0:offsetNum]
	strB := pwd[offsetNum+2:]
	return strF + strB
}

// Md5 加密
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
