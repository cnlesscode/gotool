package greflect

import (
	"fmt"
	"reflect"
	"strings"
)

// 获取变量类型
func GetType(structPoint interface{}) string {
	return reflect.TypeOf(structPoint).String()
}

// 获取结构体名称
func GetStructName(structPoint interface{}) string {
	typeName := reflect.TypeOf(structPoint).String()
	typeNameSlice := strings.Split(typeName, ".")
	return typeNameSlice[len(typeNameSlice)-1]
}

// 获取数据主键值
func PrimaryKeyValue(structPoint interface{}, PrimaryKey string) int {
	affectedId := 0
	reflectValue := reflect.ValueOf(structPoint).Elem()
	res := reflectValue.FieldByName(PrimaryKey)
	if res.Kind() == reflect.Int {
		affectedId = int(res.Int())
	}
	return affectedId
}

// 打印结构体变量
func PrintStructVars(structPoint interface{}) {
	items := reflect.ValueOf(structPoint).Elem()
	typeOfST := items.Type()
	for i := 0; i < items.NumField(); i++ {
		itemInFor := items.Field(i)
		fmt.Printf("%v : %v\n", typeOfST.Field(i).Name, itemInFor.Interface())
	}
}

// 打印结构体方法
func PrintStructMethods(structPoint interface{}) {
	items := reflect.ValueOf(structPoint).Elem()
	typeOfST := items.Type()
	for i := 0; i < items.NumMethod(); i++ {
		fmt.Printf("%v\n", typeOfST.Method(i).Name)
	}
}

// 打印结构体标签
func PrintStructTags(structPoint interface{}) {
	t := reflect.TypeOf(structPoint).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag)
	}
}
