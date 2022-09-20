### 反射工具包

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/greflect"
)
```

#### GetType(structPoint interface{}) string
函数功能 : 获取变量类型
返回格式 : string
```go
var str string = "abc"
func main() {
	stName := greflect.GetType(str)
	fmt.Printf("%v\n", stName)
}
```

#### GetStructName(structPoint interface{}) string
函数功能 : 获取变量对应结构体名称
返回格式 : string
```go
type TestStruct struct {
	Id   int
	Name string 
}

func main() {
	stName := greflect.GetStructName(TestStruct{})
	fmt.Printf("stName: %v\n", stName)
}
```

#### PrintStructVars(structPoint interface{})
函数功能 : 打印结构体变量
```go
type TestStruct struct {
	Id   int    `gorm:"id;primaryKey"`
	Name string `gorm:"name"`
}

func main() {
	greflect.PrintStructVars(&TestStruct{Id: 1, Name: "test"})
}
```

#### PrintStructVars(structPoint interface{})
函数功能 : 打印结构体变量
```go
type TestStruct struct {
	Id   int    `gorm:"id;primaryKey"`
	Name string `gorm:"name"`
}

func main() {
	greflect.PrintStructVars(&TestStruct{Id: 1, Name: "test"})
}
```


#### PrintStructMethods(structPoint interface{})
函数功能 : 打印结构体方法
```go
type TestStruct struct {
	Id   int    `gorm:"id;primaryKey"`
	Name string `gorm:"name"`
}

func (st TestStruct) Say() {
	println("say ..." + st.Name)
}

func main() {
	st := &TestStruct{}
	greflect.PrintStructMethods(st)
}
```

#### PrintStructTags(structPoint interface{})
函数功能 : 打印结构体标签
```go
type TestStruct struct {
	Id   int    `gorm:"id;primaryKey"`
	Name string `gorm:"name"`
}

func main() {
	st := TestStruct{}
	greflect.PrintStructTags(&st)
}
```

