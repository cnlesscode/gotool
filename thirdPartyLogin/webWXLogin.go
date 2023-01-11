package thirdPartyLogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cnlesscode/gotool/config"
	"github.com/cnlesscode/gotool/gmd5"
	"github.com/cnlesscode/gotool/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var WebWXLoginSets map[string]string

func init() {
	WebWXLoginSets = make(map[string]string)
	WebWXLoginSets["AppId"] = config.Ini.Section("WebWXLogin").Key("AppId").String()
	WebWXLoginSets["AppSecret"] = config.Ini.Section("WebWXLogin").Key("AppSecret").String()
	WebWXLoginSets["RedirectURI"] = config.Ini.Section("WebWXLogin").Key("RedirectURI").String()
	WebWXLoginSets["StatePrefix"] = config.Ini.Section("WebWXLogin").Key("StatePrefix").String()
}

type WebWXToken struct {
	AccessToken string `json:"access_token"`
	Openid      string `json:"openid"`
	Unionid     string `json:"unionid"`
	Errcode     int    `json:"errcode"`
}

type WebWXUserInfo struct {
	HeadImgUrl string `json:"headimgurl"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
	Nickname   string `json:"nickname"`
	Errcode    int    `json:"errcode"`
}

// 授权后返回处理
func WebWXLoginBack(ctx *gin.Context, session sessions.Session) (WebWXUserInfo, error) {
	webWXUserInfo := WebWXUserInfo{}
	// 检查返回数据
	code := ctx.Query("code")
	if code == "" {
		return webWXUserInfo, errors.New("登陆数据错误")
	}
	state := ctx.Query("state")
	if state == "" {
		return webWXUserInfo, errors.New("登陆数据错误")
	}
	// 检查 state
	stateInSession := session.Get("WebWXLoginStateCode")
	if stateInSession != state {
		return webWXUserInfo, errors.New("登陆状态数据错误")
	}
	// 通过 code 获取access_token
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("appid", WebWXLoginSets["AppId"])
	params.Add("secret", WebWXLoginSets["AppSecret"])
	params.Add("code", code)
	params.Add("redirect_uri", WebQQLoginSets["RedirectURI"])
	loginURL := fmt.Sprintf("%s?%s", "https://api.weixin.qq.com/sns/oauth2/access_token", params.Encode())
	response, err := http.Get(loginURL)
	if err != nil {
		return webWXUserInfo, err
	}
	defer response.Body.Close()
	bs, _ := ioutil.ReadAll(response.Body)
	// 成功
	// {
	// "access_token":"ACCESS_TOKEN",
	// "expires_in":7200,
	// "refresh_token":"REFRESH_TOKEN",
	// "openid":"OPENID",
	// "scope":"SCOPE",
	// "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
	// }
	// 失败 {"errcode":40001, "errmsg":"***" }
	webWXToken := &WebWXToken{}
	err = json.Unmarshal(bs, webWXToken)
	if err != nil {
		return webWXUserInfo, err
	}
	if webWXToken.Errcode > 0 {
		return webWXUserInfo, errors.New("登录失败 Error_100006")
	}
	fmt.Printf("webWXToken: %v\n", webWXToken)

	// 通过access_token调用接口 获取用户个人信息
	params2 := url.Values{}
	params2.Add("access_token", webWXToken.AccessToken)
	params2.Add("openid", webWXToken.Openid)
	loginURL = fmt.Sprintf("%s?%s", "https://api.weixin.qq.com/sns/userinfo", params2.Encode())
	fmt.Printf("loginURL: %v\n", loginURL)
	// https://api.weixin.qq.com/sns/userinfo?access_token=***&openid=***
	response2, err := http.Get(loginURL)
	if err != nil {
		return webWXUserInfo, err
	}
	defer response2.Body.Close()
	bs2, _ := ioutil.ReadAll(response2.Body)
	// 成功 {
	// "openid":"***",
	// "nickname":"***",
	// "sex":0,
	// "headimgurl":"http...",
	// "unionid":"******"
	//}
	// 失败 {"errcode":40001, "errmsg":"***" }
	err = json.Unmarshal(bs2, &webWXUserInfo)
	if err != nil {
		return webWXUserInfo, err
	}
	if webWXUserInfo.Errcode > 0 {
		return webWXUserInfo, errors.New("登录失败 Error_100008")
	}
	fmt.Printf("webWXUserInfo: %v\n", webWXUserInfo)
	return webWXUserInfo, nil
}

// 跳转到微信扫码登陆
func WebWXLogin(ctx *gin.Context, session sessions.Session) {
	// 组合 url
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("appid", WebWXLoginSets["AppId"])
	params.Add("scope", "snsapi_login")
	// 生成随机码
	randCode := gmd5.Md5(WebWXLoginSets["StatePrefix"] + random.UUID())
	// 利用 session 记录 state
	session.Set("WebWXLoginStateCode", randCode)
	session.Save()
	params.Add("state", randCode)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), WebWXLoginSets["RedirectURI"])
	loginURL := fmt.Sprintf("%s?%s", "https://open.weixin.qq.com/connect/qrconnect", str)
	// 重定向到 QQ 互联
	ctx.Redirect(302, loginURL)
}
