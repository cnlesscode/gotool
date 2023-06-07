package gstring

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"
)

// 首字母大写
func FirstUpper(str string) string {
	if str == "" {
		return ""
	}
	str = strings.ToLower(str)
	return strings.ToUpper(str[:1]) + str[1:]
}

// 首字母大写 多个单词
func FirstUpperAll(str string, removeSpace bool) string {
	reg := regexp.MustCompile("[ _]")
	strs := reg.Split(str, -1)
	for k, v := range strs {
		strs[k] = FirstUpper(strings.ToLower(v))
	}
	if removeSpace {
		return strings.Join(strs, "")
	} else {
		return strings.Join(strs, " ")
	}
}

// 首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 全部小写
func LowerAll(str string, removeSpace bool) string {
	reg := regexp.MustCompile("[ _]")
	strs := reg.Split(str, -1)
	for k, v := range strs {
		strs[k] = strings.ToLower(v)
	}
	if removeSpace {
		return strings.Join(strs, "")
	} else {
		return strings.Join(strs, " ")
	}
}

// Find Images From Html
func FindImagesFromHtml(html string) []string {
	images := make([]string, 0)
	reg := regexp.MustCompile(`(?U)<img src="(.*)"`)
	res := reg.FindAllStringSubmatch(html, -1)
	for _, v := range res {
		images = append(images, v[1])
	}
	return images
}

// Trim html tags
func TrimHtmlTags(html string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	html = re.ReplaceAllString(html, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	html = re.ReplaceAllString(html, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	html = re.ReplaceAllString(html, "")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	html = re.ReplaceAllString(html, "")
	//&..;
	re, _ = regexp.Compile("(?U)&.*;")
	html = re.ReplaceAllString(html, "")
	return strings.TrimSpace(html)
}

// Any to string
func AnyToString(value interface{}) string {
	var key string
	if value == nil {
		return ""
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

// 去除标点
func StripPunctuation(str string, replaceUnderLine bool) string {
	var reg *regexp.Regexp
	var err error
	if replaceUnderLine {
		reg, err = regexp.Compile("[\\s\\n\\t\\-,\\.\\?\\!\"@，。'、？！\\:：“”；()（）；{}【】_《》~\\*\\<\\>/\\|\\-\\+\\=\\&\\^\\%\\#\\`\\;$￥‘’〉〈…＞＜＠＃＄％︿＆＊＋～｜［］·｛｝,\\[\\]]")
	} else {
		reg, err = regexp.Compile("[\\s\\n\\t\\-,\\.\\?\\!\"@，。'、？！\\:：“”；()（）；{}【】《》~\\*\\<\\>/\\|\\-\\+\\=\\&\\^\\%\\#\\`\\;$￥‘’〉〈…＞＜＠＃＄％︿＆＊＋～｜［］·｛｝,\\[\\]]")
	}
	if err == nil {
		return reg.ReplaceAllString(str, "")
	}
	return str
}

// 格式化空格
func FormatSpace(str string) string {
	str = strings.Trim(str, " ")
	// 将连续空格替换为一个
	reg, _ := regexp.Compile(" {2,}")
	return reg.ReplaceAllString(str, " ")
}

// 使用空格将字符串拆分为切片
func StringToSliceBySpace(str string) []string {
	return strings.Split(FormatSpace(str), " ")
}
