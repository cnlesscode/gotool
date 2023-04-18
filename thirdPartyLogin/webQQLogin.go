package thirdPartyLogin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cnlesscode/gotool/config"
	"github.com/cnlesscode/gotool/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var WebQQLoginSets map[string]string

func init() {
	WebQQLoginSets = make(map[string]string)
	WebQQLoginSets["AppId"] = config.Ini.Section("WebQQLogin").Key("AppId").String()
	WebQQLoginSets["AppKey"] = config.Ini.Section("WebQQLogin").Key("AppKey").String()
	WebQQLoginSets["RedirectURI"] = config.Ini.Section("WebQQLogin").Key("RedirectURI").String()
	WebQQLoginSets["StatePrefix"] = config.Ini.Section("WebQQLogin").Key("StatePrefix").String()
}

// 解析返回数据
func WebQQLoginGetUser(ctx *gin.Context, session sessions.Session) (map[string]any, error) {
	// 检查返回数据
	code := ctx.Query("code")
	if code == "" {
		return nil, errors.New("登陆数据错误")
	}
	state := ctx.Query("state")
	if state == "" {
		return nil, errors.New("登陆数据错误")
	}

	// 检查 state
	stateInSession := session.Get("WebQQLoginStateCode")
	if stateInSession != state {
		return nil, errors.New("登陆状态数据错误")
	}

	// 获取 token
	token, err := WebQQLoginGetToken(code)
	if err != nil {
		return nil, err
	}

	// 获取 openid
	openid, err := WebQQLoginGetOpenId(token["access_token"], code)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	params := url.Values{}
	params.Add("access_token", token["access_token"])
	params.Add("openid", openid)
	params.Add("oauth_consumer_key", WebQQLoginSets["AppId"])
	uri := fmt.Sprintf("https://graph.qq.com/user/get_user_info?%s", params.Encode())
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	mapUser := make(map[string]any)
	err = json.Unmarshal(bs, &mapUser)
	if err != nil {
		return nil, err
	}
	mapUser["openid"] = openid

	return mapUser, nil
}

// 获取 toen
func WebQQLoginGetToken(code string) (map[string]string, error) {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", WebQQLoginSets["AppId"])
	params.Add("client_secret", WebQQLoginSets["AppKey"])
	params.Add("code", code)
	params.Add("redirect_uri", WebQQLoginSets["RedirectURI"])
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/token", params.Encode())
	response, err := http.Get(loginURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bs, _ := io.ReadAll(response.Body)
	body := string(bs)
	if strings.Contains(body, "access_token") {
		resultMap := ConvertToMap(body)
		return resultMap, nil
	}
	return nil, errors.New("令牌获取失败")
}

// 获取 openId
func WebQQLoginGetOpenId(token string, code string) (string, error) {
	resp, err := http.Get("https://graph.qq.com/oauth2.0/me?access_token=" + token)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, _ := io.ReadAll(resp.Body)
	body := string(bs)
	// callback( {"client_id":"YOUR_APPID","openid":"YOUR_OPENID"} );
	body = strings.Replace(body, "callback(", "", -1)
	body = strings.Replace(body, ");", "", -1)
	mapOpenId := make(map[string]string)
	err = json.Unmarshal([]byte(body), &mapOpenId)
	if err != nil {
		return "", err
	}
	return mapOpenId["openid"], nil
}

// 跳转到 QQ 登陆
func WebQQLogin(ctx *gin.Context, session sessions.Session) {
	// 组合 url
	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", WebQQLoginSets["AppId"])
	// 生成随机码
	randCode := WebQQLoginSets["StatePrefix"] + random.UUID()
	// 利用 session 记录 state
	session.Set("WebQQLoginStateCode", randCode)
	session.Save()
	params.Add("state", randCode)
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), WebQQLoginSets["RedirectURI"])
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/authorize", str)
	// 重定向到 QQ 互联
	ctx.Redirect(302, loginURL)
}

// url 转 map
func ConvertToMap(str string) map[string]string {
	var resultMap = make(map[string]string)
	values := strings.Split(str, "&")
	for _, value := range values {
		vs := strings.Split(value, "=")
		resultMap[vs[0]] = vs[1]
	}
	return resultMap
}
