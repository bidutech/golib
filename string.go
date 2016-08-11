package SHcommon

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"sync"
)

func test() {
	fmt.Println("")
}

/*
统计url保证参数唯一性记录[path][param][value1,value2,value3]
*/
type MultiMapStrings struct { //Data[key][subkey]=[]string{}
	Data map[string]map[string][]string
	sync.RWMutex
}

func NewMultiMapStrings() *MultiMapStrings {
	return &MultiMapStrings{
		Data: map[string]map[string][]string{},
	}
}

func (u *MultiMapStrings) JsonString() (str string, er error) {
	u.Lock()
	defer u.Unlock()
	b, err := json.Marshal(u.Data)
	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func (u *MultiMapStrings) Set(key, subkey, value string) {
	u.Lock()
	defer u.Unlock()
	if v, exist := u.Data[key]; exist {
		if v, exist := v[subkey]; exist {
			if FindString(v, value) {
				return
			} else {
				u.Data[key][subkey] = append(u.Data[key][subkey], value)
			}
		} else {
			u.Data[key][subkey] = []string{}
			u.Data[key][subkey] = append(u.Data[key][subkey], value)
		}

	} else {
		u.Data[key] = map[string][]string{}
		u.Data[key][subkey] = []string{}
		u.Data[key][subkey] = append(u.Data[key][subkey], value)
	}
}

func (u *MultiMapStrings) Len(key, subkey string) (length int) {
	len := len(u.Data[key][subkey])
	return len
}

/***********************************************
 如果 string 数组含有 str （完全匹配）返回true
***********************************************/
func FindString(strs []string, str string) bool {
	sort.Strings(strs)
	index := sort.SearchStrings(strs, str)
	if index < 0 || index > (len(strs)-1) {
		return false
	}
	if strs[index] == str {
		return true
	}
	return false
}

/************************
func Substr(str string, start, length int) string
对str截取 start开始长度为length的子串
****************************/
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

/**
使用 splitstr 分割 s ，找splitstr在字符串s中最后的位置
并以此位置进行分割为front和 after两部分 并返回 splitstr所在位置index
如果没有找到splitstr 则 index 为 -1

**/
func StrSplit(s string, splitstr string) (front, after string, index int) {

	n := strings.Index(s, splitstr)
	index = n
	if n < 0 {
		front = s
		after = ""
		return
	}
	front = Substr(s, 0, n)
	after = Substr(s, n, len(s))
	return
}

/**
使用 splitstr 分割 s ，找splitstr在字符串s中最后的位置
并以此位置进行分割为front和 after两部分 并返回 splitstr所在位置index
如果没有找到splitstr 则 index 为 -1

**/
func StrLastSplit(s string, splitstr string) (front, after string, index int) {

	n := strings.LastIndex(s, splitstr)
	index = n
	if n < 0 {
		front = s
		after = ""
		return
	}
	front = Substr(s, 0, n)
	after = Substr(s, n, len(s))
	return
}
