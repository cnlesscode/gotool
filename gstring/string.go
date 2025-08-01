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
	re, _ := regexp.Compile(`<[\S\s]+?>`)
	html = re.ReplaceAllStringFunc(html, strings.ToLower)
	// ...
	re, _ = regexp.Compile(`<style[\S\s]+?</style>`)
	html = re.ReplaceAllString(html, "")
	// ...
	re, _ = regexp.Compile(`<script[\S\s]+?</script>`)
	html = re.ReplaceAllString(html, "")
	// ...
	re, _ = regexp.Compile(`<[\S\s]+?>`)
	html = re.ReplaceAllString(html, "")
	// ...
	re, _ = regexp.Compile(`\s{2,}`)
	html = re.ReplaceAllString(html, "")
	// ...
	re, _ = regexp.Compile(`(?U)&.*;`)
	html = re.ReplaceAllString(html, "")
	return strings.TrimSpace(html)
}

// Any to string
func AnyToString(value interface{}) string {
	var key string
	if value == nil {
		return ""
	}
	switch value := value.(type) {
	case float64:
		key = strconv.FormatFloat(value, 'f', -1, 64)
	case float32:
		key = strconv.FormatFloat(float64(value), 'f', -1, 64)
	case int:
		key = strconv.Itoa(value)
	case uint:
		key = strconv.Itoa(int(value))
	case int8:
		key = strconv.Itoa(int(value))
	case uint8:
		key = strconv.Itoa(int(value))
	case int16:
		key = strconv.Itoa(int(value))
	case uint16:
		key = strconv.Itoa(int(value))
	case int32:
		key = strconv.Itoa(int(value))
	case uint32:
		key = strconv.Itoa(int(value))
	case int64:
		key = strconv.FormatInt(value, 10)
	case uint64:
		key = strconv.FormatUint(value, 10)
	case string:
		key = value
	case []byte:
		key = string(value)
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
