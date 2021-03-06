package siteauth

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/astaxie/beego/validation"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/siteauth/models"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
)

// SiteAuthController struct
type SiteAuthController struct {
	controllers.BaseController
}

// TemplageBaseURL 模板路径
var TemplageBaseURL = "siteauth"

// BaseTitle 页面标题
var BaseTitle = "登陆"

// Login func
func (e *SiteAuthController) Login() {

	e.Data["URL_CheckLogin"] = fmt.Sprintf("/%s/%s/checklogin", admin.AdminBaseRoterVersion, TemplageBaseURL)

	// 获取请求的 Referer ，确保登录失效以后可以返回到原来的页面
	if e.GetString("Referer") != "" {
		e.Data["Referer"] = e.GetString("Referer")
	} else {
		e.Data["Referer"] = fmt.Sprintf("/%s%s", admin.AdminBaseRoterVersion, utils.TConfig.String("WebSite::URL_AdmonHomePage"))
	}
	e.Data["title"] = fmt.Sprintf("%s管理", BaseTitle)

	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts([]string{
		//jquery-migrate
		// "/static/js/lib/jquery/jquery-migrate/1.4.1/jquery-migrate-1.4.1.min.js",
		//jquery.device.detector
		"/static/js/lib/jquery/jquery.device.detector/1.0.0/jquery.device.detector.min.js",
		//layer
		"/static/js/hplus/plugins/layer/layer.min.js",
		//sweetalert
		"/static/js/hplus/plugins/sweetalert/2.1.0/sweetalert.min.js", //第三方插件
		//recaptcha
		"https://www.google.com/recaptcha/api.js",
	})
	e.Layout = "common/page/layout-home.html"
	// e.LayoutSections = make(map[string]string)
	// e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/login.html", TemplageBaseURL)
}

// Logout 退出登录
func (e *SiteAuthController) Logout() {

	e.Ctx.SetCookie("token", "", 0)
	e.Ctx.Redirect(302, fmt.Sprintf("/%s%s", admin.AdminBaseRoterVersion, utils.TConfig.String("WebSite::URL_Login")))
}

// CheckLogin func
func (e *SiteAuthController) CheckLogin() {

	var err error
	var result models.Result
	valid := validation.Validation{}

	username := e.GetString("username")
	password := e.GetString("password")
	deviceDetector := e.GetString("device_detector")
	geoRemoteAddr := e.GetString("geo_remote_addr")
	geoCountry := e.GetString("geo_country")
	geoCity := e.GetString("geo_city")

	if !valid.MinSize(password, 3, "minSize1").Ok {
		result.Code = -1
		result.Msg = "密码长度不能小于3"
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	if username == "" || password == "" {
		result.Code = -1
		result.Msg = "用户名和密码不能为空"
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("UserName", username)
	params.Set("PassWord", password)
	params.Set("DeviceDetector", deviceDetector)
	params.Set("GeoRemoteAddr", geoRemoteAddr)
	params.Set("GeoCountry", geoCountry)
	params.Set("GeoCity", geoCity)

	// 临时 Json解析类
	responseJSON := struct {
		TokenString string `json:"tokenstring"`
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_User_UserLogin", bytes.NewBufferString(params.Encode()), &responseJSON, true, false, false, false); err != nil {
		// debug
		if utils.RunMode == "dev" {
			inner.Mlogger.Infof("登录失败：%s", err.Error())
		}

		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
	} else {
		// debug
		if utils.RunMode == "dev" {
			inner.Mlogger.Infof("登录成功：%s", responseJSON.TokenString)
		}
		// 写入 cookie,10分钟后过期
		e.Ctx.Output.Cookie("mshk_token", responseJSON.TokenString, 60*10)
		result.Code = 0
		result.Msg = "登录成功"
		e.Data["json"] = result
		e.ServeJSON()
	}

	// m, err := new(models.User).QueryOne(orm.NewCondition().And("username", username))
	// if err != nil {
	// 	if commonutils.StringContains(err.Error(), "no row found") {
	// 		result.Code = -1
	// 		result.Msg = "用户不存在"
	// 		e.Data["json"] = result
	// 		e.ServeJSON()
	// 		return
	// 	} else {
	// 		result.Code = -1
	// 		result.Msg = err.Error()
	// 		e.Data["json"] = result
	// 		e.ServeJSON()
	// 		return
	// 	}
	// }

	// if commonutils.HexEncodeToString(commonutils.GetHMAC(commonutils.HashSHA256, []byte(password), []byte(m.AuthKey))) != m.PassWord {
	// 	result.Code = -1
	// 	result.Msg = "密码不正确"
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// e.SetSession(internal.SESSIONKEYUSERID, m.ID)
	// e.SetSession(internal.SESSIONKEYUSERNAME, m.UserName)
	// //url.QueryEscape(e.Ctx.Request.RequestURI)

	// //记录登录历史
	// historyUserLogin := &models.HistoryUserLogin{
	// 	User:           &models.User{ID: m.ID},
	// 	DeviceDetector: deviceDetector,
	// 	GeoRemoteAddr:  geoRemoteAddr,
	// 	GeoCountry:     geoCountry,
	// 	GeoCity:        geoCity,
	// 	AddTime:        time.Now(),
	// }
	// if _, err := historyUserLogin.Add(historyUserLogin); err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// if e.GetString("Referer") != "" {
	// 	url, _ := url.QueryUnescape(e.GetString("Referer"))
	// 	// e.Ctx.Redirect(302, url)
	// 	result.Code = 0
	// 	result.Msg = url
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// } else {
	// 	result.Code = 0
	// 	result.Msg = beego.AppConfig.String("Site::AdminURL")
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// 	// e.Ctx.Redirect(302, beego.AppConfig.String("Site::AdminURL"))
	// }
}
