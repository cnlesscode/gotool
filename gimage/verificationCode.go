package gimage

// https://github.com/fogleman/gg

import (
	"math/rand"

	"github.com/cnlesscode/gotool/gfs"
	"github.com/cnlesscode/gotool/random"
	"github.com/fogleman/gg"
)

type VerificationCodeConfig struct {
	Width        int
	Height       int
	CodeLength   int
	NumberLength int
	FontSize     float64
	FontPath     string
	LineNumber   int
	TmpDir       string
	CodeColor    []float64
}

func VerificationCode(conf VerificationCodeConfig) (string, string, error) {
	VCode := random.RandomCharacters(conf.CodeLength, conf.NumberLength)
	dc := gg.NewContext(conf.Width, conf.Height)

	// 绘制验证码
	dc.SetRGB(conf.CodeColor[0], conf.CodeColor[1], conf.CodeColor[2])
	err := dc.LoadFontFace(conf.FontPath, conf.FontSize)
	if err == nil {
		dc.DrawStringAnchored(VCode, float64(conf.Width)/2, float64(conf.Height)/2, 0.5, 0.5)
	} else {
		return "", "", err
	}
	// 绘制干扰线
	for i := 0; i < conf.LineNumber; i++ {
		x1, y1 := RandPos(conf.Width, conf.Height)
		x2, y2 := RandPos(conf.Width, conf.Height)
		r, g, b, a := RandColor(255)
		w := float64(rand.Intn(3) + 1)
		dc.SetRGBA255(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	// 保存图片
	if !gfs.DirExists(conf.TmpDir) {
		gfs.MakeDir(conf.TmpDir)
	}
	imageFilePath := conf.TmpDir + "/" + random.Md5UUID() + ".png"
	dc.SavePNG(imageFilePath)

	// 转为 base64
	base64data, _ := ImageToBase64(imageFilePath, true)

	// 返回
	return VCode, base64data, nil
}

// 随机坐标
func RandPos(width, height int) (x float64, y float64) {
	x = rand.Float64() * float64(width)
	y = rand.Float64() * float64(height)
	return x, y
}

// 随机颜色
func RandColor(maxColor int) (r, g, b, a int) {
	r = int(uint8(rand.Intn(maxColor)))
	g = int(uint8(rand.Intn(maxColor)))
	b = int(uint8(rand.Intn(maxColor)))
	a = int(uint8(rand.Intn(255)))
	return r, g, b, a
}
