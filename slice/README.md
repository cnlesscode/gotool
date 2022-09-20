### 切片操作工具

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/slice"
)
```

#### SortRandomlyString(字符串切片 []string) []string
函数功能 : 将字符串切片进行随机排序
返回格式 : []string
```go
slice1 := []string{"a", "b", "c", "d"}
slice1 = slice.SortRandomlyString(slice1)
fmt.Printf("slice1: %v\n", slice1)
```

#### SortRandomlyString(整数 []int) []int
函数功能 : 将整数切片进行随机排序
返回格式 : []int
```go
slice1 := []int{1, 2, 3, 4}
slice1 = slice.SortRandomlyInt(slice1)
fmt.Printf("slice1: %v\n", slice1)
```

#### SortRandomly([]interface{}) []interface{}
函数功能 : 将任意切面进行随机排序
返回格式 : []int
```go
slice1 := []interface{}{1, "hi", 3, 4}
slice1 = slice.SortRandomly(slice1)
fmt.Printf("slice1: %v\n", slice1)
```

### 切片随机排序扩展
请参考 SortRandomly 封装自己的切片排序函数（ 针对自己的切片类型 ）
```go
func SortRandomly(slice []interface{}) []interface{} {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })
	return slice
}
```

#### RemoveInt([]int, index int) []int
函数功能 : 根据指定索引删除切片元素 []int
返回格式 : []int
```go
slice1 := []int{1, 2, 3, 4}
slice1 = slice.RemoveInt(slice1, 2)
fmt.Printf("slice1: %v\n", slice1)
// slice1: [1 2 4]
```

#### RemoveInterface([]interface{}, index int) []interface{}
函数功能 : 根据指定索引删除切片元素 []interface
返回格式 : []interface{}
```go
slice1 := []interface{}{1, "hi", 3, 4}
slice1 = slice.RemoveInterface(slice1, 1)
fmt.Printf("slice1: %v\n", slice1)
// slice1: [1 3 4]
```

### 切片删除扩展
请参考 RemoveInterface 封装自己的切片删除函数（ 针对自己的切片类型 ）
```go
func RemoveInterface(slice []interface{}, index int) []interface{} {
	slice = append(slice[:index], slice[index+1:]...)
	return slice
}
```

#### AppendInt([]int, index int) []int
函数功能 : 向指定位置添加元素 []int
返回格式 : []int
```go
slice1 := []int{1, 2, 3, 4}
slice1 = slice.AppendInt(slice1, 2, 8)
fmt.Printf("slice1: %v\n", slice1)
// slice1: [1 2 8 3 4]
```

#### AppendInterFace([]interface{}, index int) []int
函数功能 : 向指定位置添加元素 []interface{}
返回格式 : []interface{}
```go
slice1 := []interface{}{1, "a", 3, 4}
slice1 = slice.AppendInterFace(slice1, 2, "test")
fmt.Printf("slice1: %v\n", slice1)
// slice1: [1 a test 3 4]
```

#### 切片添加元素扩展
函数功能 : 查找指定元素 []interface{}
返回格式 : int 
```go
slice1 := []interface{}{1, "hi", 3, 4}
idx := slice.Find(slice1, "hi")
fmt.Printf("idx: %v\n", idx)
```

#### 查找元素
请参考 AppendInterFace 封装自己的切片元素添加函数（ 针对自己的切片类型 ）
```go
func AppendInterFace(slice []interface{}, index int, val interface{}) []interface{} {
	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = val
	return slice
}
```

#### 切片交集与差集
在实际的开发过程中切片类型多种多样，我们提供了整数类型的切片交集、差集算法，请根据自己的类型改进使用 ~
```go
package main

import "fmt"

// 切片差集
func Difference(slice1, slice2 []int) []int {

	m := make(map[int]int)
	n := make([]int, 0)
	inter := Intersect(slice1, slice2)
	// 获得交集
	for _, v := range inter {
		m[v] = 1
	}
	// 变量第1个切片, 不是交集则为差集
	for _, value := range slice1 {
		if m[value] == 0 {
			n = append(n, value)
		}
	}
	// 变量第2个切片, 不是交集则为差集
	for _, v := range slice2 {
		if m[v] == 0 {
			n = append(n, v)
		}
	}
	return n
}

// 两个切片的交集
func Intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	n := make([]int, 0)
	for _, v := range slice1 {
		m[v] = 1
	}
	fmt.Printf("m: %v\n", m)
	for _, v := range slice2 {
		times, ok := m[v]
		if ok && times == 1 {
			n = append(n, v)
		}
	}
	return n
}

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{2, 3, 4}
	// 交集
	intersect := Intersect(slice1, slice2)
	fmt.Printf("intersect: %v\n", intersect)
	// 差集
	difference := Difference(slice1, slice2)
	fmt.Printf("difference: %v\n", difference)
}

```

### ToStringItems
函数功能 : 将任意类型切片换为字符串类型切片
返回格式 : []string
```go
package main
import (
	"fmt"
	"github.com/cnlesscode/gotool/slice"
)
func main() {
	a := []interface{}{1, 2, "a", "b"}
	aString := slice.ToStringItems(a)
	fmt.Printf("aString: %v\n", aString)
}
```

### ToString
函数功能 : 将任意切片转换为字符串
返回格式 : string
```go
func main() {
	a := []interface{}{1, 2, "a", "b"}
	aString := slice.ToString(a, "-")
	fmt.Printf("aString: %v\n", aString)
}
```