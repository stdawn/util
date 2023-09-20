/*
@Time : 2022/5/20 13:32
@Author : LiuKun
@File : tools
@Software: GoLand
@Description:
*/

package util

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"
)

// RemoveNonNumeric 去除非数字部分
//输入"abc123def456"， 输出 "123456"
func RemoveNonNumeric(str string) string {
	regex := regexp.MustCompile("[^0-9]+")
	result := regex.ReplaceAllString(str, "")
	return result
}

// GetDateFromString 从字符串获取日期，返回年月日
func GetDateFromString(s string) (int, int, int) {
	year, month, day := 0, 0, 0
	temp := s
	if strings.Contains(temp, "年") {
		ss := strings.Split(temp, "年")
		year = GetIntFromV(ss[0])
		if len(ss) > 1 {
			temp = ss[1]
		}
	}
	if strings.Contains(temp, "月") {
		ss := strings.Split(temp, "月")
		month = GetIntFromV(ss[0])
		if len(ss) > 1 {
			temp = ss[1]
		}
	}
	if strings.Contains(temp, "日") {
		ss := strings.Split(temp, "日")
		day = GetIntFromV(ss[0])
	}

	return year, month, day
}

// DurationDes 时长描述
func DurationDes(d time.Duration) string {
	s := ""
	hour := int64(d) / int64(time.Hour)
	if hour > 0 {
		s += fmt.Sprintf("%d小时", hour)
	}
	minute := (int64(d) % int64(time.Hour)) / int64(time.Minute)
	if minute > 0 {
		s += fmt.Sprintf("%d分", minute)
	}
	second := float64((int64(d)%int64(time.Hour))%int64(time.Minute)) / float64(time.Second)
	if second > 0 {
		s += fmt.Sprintf("%.3f秒", second)
	}
	return s
}

// LetterMarkSecondLastChinese 字母标记倒数第二个汉字
func LetterMarkSecondLastChinese(str string) string {
	indexes := make([]int, 0)
	rt := []rune(str)
	for i, v := range rt {
		if unicode.Is(unicode.Han, v) {
			indexes = append(indexes, i)
		}
	}
	if len(indexes) >= 2 {
		oStr := string(rt[indexes[len(indexes)-2]])

		pins := pinyin.Pinyin(oStr, pinyin.NewArgs())
		if len(pins) > 0 && len(pins[0]) > 0 {
			p := pins[0][0]
			if len(p) > 0 {
				rStr := strings.ToUpper(p[:1])

				count := strings.Count(str, oStr)
				if count < 2 {
					//只出现一个
					return strings.ReplaceAll(str, oStr, rStr)
				}

				//步步高、拼多多、新论新材
				replaceCount := count - 1
				lastStr := string(rt[indexes[len(indexes)-1]])
				if lastStr == oStr {
					replaceCount -= 1
				}
				newStr := strings.Replace(str, oStr, "ShouldReplace", replaceCount)
				newStr = strings.Replace(newStr, oStr, rStr, 1)
				newStr = strings.ReplaceAll(newStr, "ShouldReplace", oStr)

				return newStr

			}
		}
	}

	return str
}

// GetMapValueSlice 获取Value值切片, 不存在的Key为空字符串
func GetMapValueSlice(keys []string, m map[string]string) []string {
	vs := make([]string, 0)
	for _, k := range keys {
		v, ok := m[k]
		if ok {
			vs = append(vs, v)
		} else {
			vs = append(vs, "")
		}
	}
	return vs
}

// RemoveDuplicatedString  去除重复的字符串
func RemoveDuplicatedString(slice []string) []string {

	ss := make([]string, 0)
	m := make(map[string]bool)
	for _, s := range slice {
		_, ok := m[s]
		if !ok {
			m[s] = true
			ss = append(ss, s)
		}
	}
	return ss
}

// MatchKeys  匹配关键字 ([index], map[index]key)
func MatchKeys(text string, keys []string) ([]int, map[int]string) {

	containsMaps := make(map[int]string)

	for _, key := range keys {
		i := strings.Index(text, key)
		if i >= 0 {
			k, ok := containsMaps[i]
			if ok {
				if len(k) > len(key) {
					//优先匹配长的，规避“东方通信” 被 “东方通” 和 “东方通讯同时匹配”
					containsMaps[i] = k
				}
			} else {
				containsMaps[i] = key
			}
		}
	}

	indexes := make([]int, 0)
	for k := range containsMaps {
		indexes = append(indexes, k)
	}

	sort.Ints(indexes)
	return indexes, containsMaps
}

// FloatSliceToString  浮点数数组转字符串
func FloatSliceToString(fs []float64, split string) string {
	ns := ""
	for i, f := range fs {
		ns += fmt.Sprintf("%.2f", f)
		if i < len(fs)-1 {
			ns += split
		}
	}
	return ns
}

// FloatPercentSliceToString  浮点百分比数组转字符串
func FloatPercentSliceToString(fs []float64, split string) string {
	ns := ""
	for i, f := range fs {
		ns += fmt.Sprintf("%.2f%%", f*100)
		if i < len(fs)-1 {
			ns += split
		}
	}
	return ns
}

// IntSliceToString  整数数组转字符串
func IntSliceToString(is []int, split string) string {
	ns := ""
	for i, f := range is {
		ns += fmt.Sprintf("%d", f)
		if i < len(is)-1 {
			ns += split
		}
	}
	return ns
}

// StringSliceToString 字符串数组转字符串
func StringSliceToString(ss []string, split string) string {
	ns := ""
	for i, s := range ss {
		ns += s
		if i < len(ss)-1 {
			ns += split
		}
	}
	return ns
}

// InterfaceSliceToString object数组转字符串
func InterfaceSliceToString(ss []interface{}, split string) string {
	ns := ""
	for i, s := range ss {
		ns += GetStringFromV(s)
		if i < len(ss)-1 {
			ns += split
		}
	}
	return ns
}
