package service

import (
	"github.com/liuhongdi/digv19/cache"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// 设置自带的store
//var store = base64Captcha.DefaultMemStore

//使用redis作为store
var store = cache.RedisStore{}

//生成验证码
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverString{
		Height:          60,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

//验证captcha是否正确
func CaptVerify(id string,capt string) bool {
	if store.Verify(id, capt, false) {
		return true
	} else {
		return false
	}

}