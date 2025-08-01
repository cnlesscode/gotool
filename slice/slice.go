package slice

import (
	"math/rand"
	"strings"
	"time"

	"github.com/cnlesscode/gotool/gstring"
)

// Sort slices randomly []string
func SortRandomlyString(slice []string) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}

func SortRandomlyInt(slice []int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}

func SortRandomly(slice []interface{}) []interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}

// Remove slice item by index []int
func RemoveInt(slice []int, index int) []int {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

// Remove slice item by index []interface
func RemoveInterface(slice []interface{}, index int) []interface{} {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}

// Append an item by index
func AppendInt(slice []int, index int, val int) []int {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = val
	return slice
}

// Append an item by index
func AppendInterFace(slice []interface{}, index int, val interface{}) []interface{} {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = val
	return slice
}

// Find item
func Find(slice []interface{}, val interface{}) int {
	index := -1
	for idx, v := range slice {
		if v == val {
			index = idx
			break
		}
	}
	return index
}

// 字符串切片差集
func DifferenceString(a, b []string) []string {
	var res []string
	for _, va := range a {
		idx := -1
		for fIdx, vb := range b {
			if va == vb {
				idx = fIdx
				break
			}
		}
		if idx == -1 {
			res = append(res, va)
		}
	}
	return res
}

// To String Items
func ToStringItems(slice []interface{}) []string {
	res := make([]string, 0)
	for _, v := range slice {
		res = append(res, gstring.AnyToString(v))
	}
	return res
}

// To String
func ToString(slice []interface{}, sep string) string {
	res := make([]string, 0)
	for _, v := range slice {
		res = append(res, gstring.AnyToString(v))
	}
	return strings.Join(res, sep)
}
