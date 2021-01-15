package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv19/pkg/result"
	"github.com/liuhongdi/digv19/service"
	//"github.com/liuhongdi/digv18/service"
)

type IdController struct{}

func NewIdController() IdController {
	return IdController{}
}

//存储验证码的结构
type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	//VerifyValue string `json:"code"`
}

// 生成图形化验证码
func (a *IdController) GetOne(ctx *gin.Context) {
	resultRes := result.NewResult(ctx)
	id, b64s, err := service.CaptMake()
	if err != nil {
		resultRes.Error(1,err.Error())
	}
	captchaResult := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
	}
	resultRes.Success(captchaResult)
    return
}

//验证captcha是否正确
func (a *IdController) Verify(c *gin.Context) {

	id := c.PostForm("id")
	capt := c.PostForm("capt")
	resultRes := result.NewResult(c)
	if (id == "" || capt == "") {
		resultRes.Error(400,"param error")
	}

	if service.CaptVerify(id, capt) == true {
		resultRes.Success("success")
	} else {
		resultRes.Error(1,"verify failed")
	}
	return
}

