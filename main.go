package main

import (
	"fmt"

	"github.com/cnlesscode/gotool/thirdPartyLogin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 开启 session
	store := cookie.NewStore([]byte("pwd..."))
	r.Use(sessions.Sessions("WESSESSION", store))

	// 跳转到微信登录页面
	r.GET("/WebWXLogin", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		thirdPartyLogin.WebWXLogin(ctx, session)
	})
	// 授权登录后返回页面
	r.GET("/WXLogin/Back", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		webWXUserInfo, err := thirdPartyLogin.WebWXLoginBack(ctx, session)
		if err != nil {
			println("登录失败 : " + err.Error())
		} else {
			fmt.Printf("res: %v\n", webWXUserInfo)
			// 此处已经获取到用户的 openid、Unionid、HeadImgUrl 等数据
			// 结构体格式如下 :
			// type WebWXUserInfo struct {
			// 	HeadImgUrl string `json:"headimgurl"`
			// 	Openid     string `json:"openid"`
			// 	Unionid    string `json:"unionid"`
			// 	Nickname   string `json:"nickname"`
			// 	Errcode    int    `json:"errcode"`
			// }
			// 利用上面的数据配合数据库，继续完成后续的用户登陆逻辑即可
		}
	})

	// 监听指定端口
	r.Run(":80")

}
