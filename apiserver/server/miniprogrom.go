package server

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/sirupsen/logrus"
)

// 客户端列表
var MiniProgramApp *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService() (*miniProgram.MiniProgram, error) {
	logrus.Infoln("miniprogram app_id: %s", )

	if MiniProgramApp !=nil{
		return MiniProgramApp,nil
	}
	MiniProgramApp, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:        MpAppid,  // 小程序、公众号或者企业微信的appid
		Secret:       MpSerect, // 商户号 appID
		ResponseType: response.TYPE_MAP,
		// Token:        conf.MiniProgram.MessageToken,
		// AESKey:       conf.MiniProgram.MessageAesKey,

		Log: miniProgram.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		HttpDebug: true,
		Debug:     false,
	})

	return MiniProgramApp, err
}

