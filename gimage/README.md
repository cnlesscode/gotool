### 图片操作工具

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/gimage"
)
```


#### Base64ToImage(data string, targetDir string, imageName string) (string, error)
函数功能 : 将 Base64 数据保存为图片
返回格式 : 图片路径, 错误
```go
imageData := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAGMAAABrCAYAAABjaDz4AAABIElEQVR4nO3csQ2AMAwAQcI6zMGi7Og6lIzAF3cTWHq5s7xmZh8knH8PwEeMEDFCxAgRI0SMEDFCxAgRI0SMEDFCxAgRI0SMEDFCxAgRI0SMEDFCxAgRI2Rdz+06JMJmhIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkaIGCFihIgRIkbImhm/0CNsRogYIWKEiBEiRogYIWKEiBEiRogYIWKEiBEiRogYIWKEiBEiRogYIWKEiBEiRogYIWKEiBHyAoHsCqFV/MBgAAAAAElFTkSuQmCC"
gimage.Base64ToImage(imageData, "./imgs/", "test")
// 返回 : ./imgs/test.png , nil
```


#### ImageToBase64(imageUri string, removeImage bool) (string, error)
函数功能 : 将图片转换为 Base64
返回格式 : base64 字符串, error
```go
imageUri := "./a.png"
base64Data := gimage.ImageToBase64(imageUri, false)
```


#### VerificationCode(conf VerificationCodeConfig) (string, string, error)
函数功能 : 绘制一个验证码码图片并返回图片的base64字符串
返回格式 : 验证码字符串, 图片base64字符串, error
字体路径 : 请下载 https://github.com/cnlesscode/gotool/tree/main/resources/font
字体部署 : 请下载字体文件并部署到您的项目根目录 /resources 文件夹下
```go
conf := gimage.VerificationCodeConfig{
	Width:        160,
	Height:       60,
	CodeLength:   6,
	NumberLength: 2,
	FontSize:     30,
	// 尺寸可以自定义字体
	FontPath:     "./resources/fonts/FiraCode.ttf",
	LineNumber:   10,
	TmpDir:       "./tmp",
	CodeColor:    []float64{0.1, 0.1, 0.1},
}
vcode, vcode64, err := gimage.VerificationCode(conf)
fmt.Printf("vcode: %v\n", vcode)
fmt.Printf("vcode64: %v\n", vcode64)
if err != nil {
	fmt.Printf("err.Error(): %v\n", err.Error())
}
```


#### 图片操作工具包推荐
<https://github.com/fogleman/gg>
